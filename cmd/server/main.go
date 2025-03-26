package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang-memory-store/internal/api"
	"golang-memory-store/internal/auth"
	"golang-memory-store/internal/core"
	"golang-memory-store/internal/persistence"

	"github.com/gorilla/mux"
)

var (
	EnablePersistence = os.Getenv("ENABLE_PERSISTENCE") == "true"
	UseDatabase       = os.Getenv("DB_TYPE") != ""
	FILE_PATH         = "data.json"
)

func initDB() {
	dbType := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DB_DSN")

	if dbType == "" {
		log.Println("No DB_TYPE specified, using File Persistence only.")
		return
	}

	err := persistence.InitDB(dsn, dbType)
	if err != nil {
		log.Fatal("Failed to initialize the database:", err)
	}
}

func main() {
	store := core.NewShardedStore()
	handler := api.NewHandler(store)

	// Initialize Database if enabled
	if UseDatabase {
		initDB()
	}

	// Load data from file or DB
	if UseDatabase {
		err := store.LoadStoreFromDB()
		if err != nil {
			log.Println("Error loading from DB:", err)
		}
	} else if EnablePersistence {
		err := store.LoadStoreFromFile(FILE_PATH)
		if err != nil {
			log.Println("No previous data found, starting fresh.")
		}
	}

	r := mux.NewRouter()

	// Authentication Route
	r.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
		}
		json.NewDecoder(r.Body).Decode(&req)

		token, err := auth.GenerateToken(req.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}).Methods("POST")

	// Protected Routes
	apiRouter := r.PathPrefix("/").Subrouter()
	apiRouter.Use(api.Authenticate)
	apiRouter.HandleFunc("/set", handler.Set).Methods("POST")
	apiRouter.HandleFunc("/get/{key}", handler.Get).Methods("GET")
	apiRouter.HandleFunc("/delete/{key}", handler.Delete).Methods("DELETE")
	apiRouter.HandleFunc("/list/push", handler.Push).Methods("POST")
	apiRouter.HandleFunc("/list/pop/{key}", handler.Pop).Methods("POST")

	// Start the server asynchronously
	go func() {
		log.Println("Server running on :8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal("Server failed:", err)
		}
	}()

	// Graceful Shutdown Handling
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	<-signalChannel // Block until an interrupt signal is received

	log.Println("Shutting down the server and saving data...")

	// Save data on shutdown
	if UseDatabase {
		store.SaveStoreToDBAsync()
	} else if EnablePersistence {
		store.SaveStoreToFileAsync(FILE_PATH)
	}

	log.Println("Data saved successfully. Goodbye!")
}
