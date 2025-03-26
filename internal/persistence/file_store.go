package persistence

import (
	"encoding/json"
	"os"
	"sync"
)

var saveMutex sync.Mutex

// SaveToFile saves the in-memory store to a JSON file.
func SaveToFile(filename string, data interface{}) error {
	saveMutex.Lock()
	defer saveMutex.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

// LoadFromFile loads the data from a JSON file to the in-memory store.
func LoadFromFile(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}
