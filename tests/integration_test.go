package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var BaseURL = "http://localhost:8080"
var Token = ""

func init() {
	// Generate a valid JWT token before running tests
	data := map[string]string{"username": "testuser"}
	body, _ := json.Marshal(data)

	resp, err := http.Post(BaseURL+"/token", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic("Unable to generate JWT token: " + err.Error())
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	Token = result["token"]

	if Token == "" {
		panic("Failed to obtain a valid token")
	}
}

func TestSetAndGet(t *testing.T) {
	data := map[string]interface{}{
		"key":   "testKey",
		"value": "testValue",
		"ttl":   60,
	}
	body, _ := json.Marshal(data)

	// Set the key
	req, _ := http.NewRequest("POST", BaseURL+"/set", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	// Get the key
	req, _ = http.NewRequest("GET", BaseURL+"/get/testKey", nil)
	req.Header.Set("Authorization", "Bearer "+Token)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}

func TestDelete(t *testing.T) {
	req, _ := http.NewRequest("DELETE", BaseURL+"/delete/testKey", nil)
	req.Header.Set("Authorization", "Bearer "+Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}
