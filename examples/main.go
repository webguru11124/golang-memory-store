package main

import (
	"fmt"
	"golang-memory-store/internal/client"
	"log"
)

func main() {
	// Initialize the client with server URL and username
	apiClient, err := client.NewClient("http://localhost:8080", "techlando")
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	// ====== Set Operation (String) ======
	err = apiClient.Set("foo", "bar", 60)
	if err != nil {
		log.Fatal("Failed to set value:", err)
	}
	fmt.Println("Set operation successful")

	// ====== Get Operation (String) ======
	value, err := apiClient.Get("foo")
	if err != nil {
		log.Fatal("Failed to get value:", err)
	}
	fmt.Printf("Get operation successful: Key='foo', Value='%v'\n", value)

	// ====== Push Operation (List) ======
	err = apiClient.Push("mylist", "item1")
	if err != nil {
		log.Fatal("Failed to push to list:", err)
	}
	err = apiClient.Push("mylist", "item2")
	if err != nil {
		log.Fatal("Failed to push to list:", err)
	}
	fmt.Println("Push operation successful")

	// ====== Pop Operation (List) ======
	item, err := apiClient.Pop("mylist")
	if err != nil {
		log.Fatal("Failed to pop from list:", err)
	}
	fmt.Printf("Pop operation successful: Popped Item='%v'\n", item)

	// ====== Delete Operation (String) ======
	err = apiClient.Delete("foo")
	if err != nil {
		log.Fatal("Failed to delete key:", err)
	}
	fmt.Println("Delete operation successful")
}
