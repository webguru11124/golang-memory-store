package core

import (
	"golang-memory-store/internal/persistence"

	"hash/fnv"
	"log"
	"sync"
	"time"
)

const ShardCount = 16

type Entry struct {
	Value      interface{}
	Expiration int64
}

type ShardedStore struct {
	shards []Store
}

type Store struct {
	data  map[string]Entry
	mutex sync.RWMutex
}

// NewShardedStore initializes a new sharded store with independent locks
func NewShardedStore() *ShardedStore {
	shards := make([]Store, ShardCount)
	for i := 0; i < ShardCount; i++ {
		shards[i] = Store{data: make(map[string]Entry)}
	}
	return &ShardedStore{shards: shards}
}

// getShard selects the shard for a given key using a hash function
func (ss *ShardedStore) getShard(key string) *Store {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return &ss.shards[int(hash.Sum32())%ShardCount]
}

// Set adds or updates a key-value pair with optional TTL (in seconds)
func (ss *ShardedStore) Set(key string, value interface{}, ttl int) {
	shard := ss.getShard(key)
	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	expiration := int64(0)
	if ttl > 0 {
		expiration = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}
	shard.data[key] = Entry{Value: value, Expiration: expiration}
}

// Get retrieves the value associated with a key
func (ss *ShardedStore) Get(key string) (interface{}, bool) {
	shard := ss.getShard(key)
	shard.mutex.RLock()
	defer shard.mutex.RUnlock()

	entry, found := shard.data[key]
	if !found || (entry.Expiration > 0 && time.Now().Unix() > entry.Expiration) {
		return nil, false
	}
	return entry.Value, true
}

// Delete removes a key-value pair from the store
func (ss *ShardedStore) Delete(key string) {
	shard := ss.getShard(key)
	shard.mutex.Lock()
	defer shard.mutex.Unlock()
	delete(shard.data, key)
}

// LoadStoreFromFile loads data from a file into the sharded store.
func (ss *ShardedStore) LoadStoreFromFile(filename string) error {
	// Read data from file
	fullData := make(map[string]Entry)
	err := persistence.LoadFromFile(filename, &fullData)
	if err != nil {
		log.Println("Error loading store from file:", err)
		return err
	}

	// Distribute data across shards
	for key, entry := range fullData {
		shard := ss.getShard(key)
		shard.mutex.Lock()
		shard.data[key] = entry
		shard.mutex.Unlock()
	}

	return nil
}

// SaveStoreToFileAsync saves the entire store to a file asynchronously
func (ss *ShardedStore) SaveStoreToFileAsync(filename string) {
	go func() {
		ss.SaveStoreToFile(filename)
	}()
}

// SaveStoreToFile saves the entire store to a file (Blocking Operation)
func (ss *ShardedStore) SaveStoreToFile(filename string) {
	fullData := make(map[string]Entry)
	for i := 0; i < ShardCount; i++ {
		shard := &ss.shards[i]
		shard.mutex.RLock()
		for key, entry := range shard.data {
			fullData[key] = entry
		}
		shard.mutex.RUnlock()
	}
	err := persistence.SaveToFile(filename, fullData)
	if err != nil {
		log.Println("Error saving store to file:", err)
	}
}

// GetList retrieves a list from the store, creating one if it doesn't exist.
func (ss *ShardedStore) GetList(key string) *List {
	shard := ss.getShard(key)
	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	entry, found := shard.data[key]
	if found {
		if list, ok := entry.Value.(*List); ok {
			return list
		}
	}

	list := NewList()
	shard.data[key] = Entry{Value: list}
	return list
}

// SaveStoreToDBAsync saves the current state of the in-memory store to the database asynchronously.
func (ss *ShardedStore) SaveStoreToDBAsync() error {
	go func() {
		ss.SaveStoreToDB()
	}()
	return nil
}

// SaveStoreToDB saves the current state of the in-memory store to the database (Blocking Operation).
func (ss *ShardedStore) SaveStoreToDB() error {
	convertedData := make(map[string]interface{})

	for i := 0; i < ShardCount; i++ {
		shard := &ss.shards[i]
		shard.mutex.RLock()
		for key, entry := range shard.data {
			convertedData[key] = entry.Value
		}
		shard.mutex.RUnlock()
	}

	err := persistence.SaveToDB(convertedData)
	if err != nil {
		log.Println("Error saving store to DB:", err)
	}
	return err
}

// LoadStoreFromDB loads data from the database into the in-memory store.
func (ss *ShardedStore) LoadStoreFromDB() error {
	data, err := persistence.LoadFromDB()
	if err != nil {
		log.Println("Error loading store from DB:", err)
		return err
	}

	for key, value := range data {
		shard := ss.getShard(key)
		shard.mutex.Lock()
		shard.data[key] = Entry{Value: value}
		shard.mutex.Unlock()
	}

	return nil
}
