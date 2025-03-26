package core

import (
	"golang-memory-store/internal/persistence"

	"log"
	"sync"
	"time"
)

type Entry struct {
	Value      interface{}
	Expiration int64
}

type Store struct {
	data  map[string]Entry
	mutex sync.RWMutex
}

const PersistenceFile = "data.json"

// NewStore initializes a new in-memory store.
func NewStore() *Store {
	return &Store{
		data: make(map[string]Entry),
	}
}

// Set adds or updates a key-value pair with optional TTL (in seconds).
func (s *Store) Set(key string, value interface{}, ttl int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	expiration := int64(0)
	if ttl > 0 {
		expiration = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	s.data[key] = Entry{Value: value, Expiration: expiration}
}

// Get retrieves the value associated with a key.
func (s *Store) Get(key string) (interface{}, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	entry, found := s.data[key]
	if !found || (entry.Expiration > 0 && time.Now().Unix() > entry.Expiration) {
		return nil, false
	}
	return entry.Value, true
}

// Delete removes a key-value pair from the store.
func (s *Store) Delete(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
}

// SaveStoreToFile saves the entire store to a file.
func (s *Store) SaveStoreToFile() error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	err := persistence.SaveToFile(PersistenceFile, s.data)
	if err != nil {
		log.Println("Error saving store to file:", err)
	}
	return err
}

// LoadStoreFromFile loads data from a file into the store.
func (s *Store) LoadStoreFromFile() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data := make(map[string]Entry)
	err := persistence.LoadFromFile(PersistenceFile, &data)
	if err != nil {
		log.Println("Error loading store from file:", err)
		return err
	}
	s.data = data
	return nil
}

// GetList retrieves a list from the store, creating one if it doesn't exist.
func (s *Store) GetList(key string) *List {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entry, found := s.data[key]
	if found {
		if list, ok := entry.Value.(*List); ok {
			return list
		}
	}

	list := NewList()
	s.data[key] = Entry{Value: list}
	return list
}

// SaveStoreToDB saves the current state of the in-memory store to the database.
func (s *Store) SaveStoreToDB() error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	convertedData := make(map[string]interface{})
	for key, entry := range s.data {
		convertedData[key] = entry.Value
	}
	err := persistence.SaveToDB(convertedData)
	if err != nil {
		log.Println("Error saving store to DB:", err)
	}
	return err
}

// LoadStoreFromDB loads data from the database into the in-memory store.
func (s *Store) LoadStoreFromDB() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := persistence.LoadFromDB()
	if err != nil {
		log.Println("Error loading store from DB:", err)
		return err
	}
	for key, value := range data {
		s.data[key] = Entry{Value: value}
	}
	return nil
}
