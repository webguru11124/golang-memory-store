package main

import (
	"fmt"
	"golang-memory-store/internal/client"
	"log"
)

func main() {
	apiClient, err := client.NewClient("http://localhost:8080", "techlando")
	if err != nil {
		log.Fatal(err)
	}

	// Set a key
	err = apiClient.Set("foo", "bar", 60)
	if err != nil {
		log.Fatal(err)
	}

	// Get the key
	value, err := apiClient.Get("foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}
