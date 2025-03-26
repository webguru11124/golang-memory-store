package core

import (
	"testing"
	"time"
)

func TestStoreSetGet(t *testing.T) {
	store := NewShardedStore()
	store.Set("key1", "value1", 0)

	val, found := store.Get("key1")
	if !found || val != "value1" {
		t.Errorf("Expected 'value1', got %v", val)
	}
}

func TestStoreExpiration(t *testing.T) {
	store := NewShardedStore()
	store.Set("tempKey", "tempValue", 1)
	time.Sleep(2 * time.Second)

	_, found := store.Get("tempKey")
	if found {
		t.Error("Expected 'tempKey' to expire")
	}
}

func TestStoreDelete(t *testing.T) {
	store := NewShardedStore()
	store.Set("deleteKey", "value", 0)
	store.Delete("deleteKey")

	_, found := store.Get("deleteKey")
	if found {
		t.Error("Expected 'deleteKey' to be deleted")
	}
}
