package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var BaseURL = "http://localhost:8080"
var Token = "YOUR_JWT_TOKEN"

func TestSetAndGet(t *testing.T) {
	data := map[string]interface{}{
		"key":   "testKey",
		"value": "testValue",
		"ttl":   60,
	}
	body, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", BaseURL+"/set", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	resp, _ = http.Get(BaseURL + "/get/testKey")
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}

func TestDelete(t *testing.T) {
	req, _ := http.NewRequest("DELETE", BaseURL+"/delete/testKey", nil)
	req.Header.Set("Authorization", "Bearer "+Token)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}
