package core

import (
	"log"
	"time"
)

var store map[string]*Obj

type Obj struct {
	Value     interface{}
	ExpiresAt int64
}

func init() {
	store = make(map[string]*Obj)
}

func NewObj(value interface{}, durationMs int64) *Obj {
	var expiresAt int64 = -1
	if durationMs > 0 {
		expiresAt = time.Now().UnixMilli() + durationMs
	}

	return &Obj{
		Value:     value,
		ExpiresAt: expiresAt,
	}
}

func Put(k string, obj *Obj) {
	store[k] = obj
}

func Get(k string) *Obj {
	return store[k]
}

func Del(k string) bool {
	if _, exists := store[k]; exists {
		delete(store, k)
		return true
	}
	return false
}

func GetAllKeys() []string {
	keys := make([]string, 0, len(store))
	for k := range store {
		keys = append(keys, k)
	}
	return keys
}

// CleanupExpiredKeys removes all expired keys from the store
func CleanupExpiredKeys() {
	now := time.Now().UnixMilli()
	keysToDelete := make([]string, 0)
	
	for key, obj := range store {
		if obj.ExpiresAt != -1 && obj.ExpiresAt <= now {
			keysToDelete = append(keysToDelete, key)
		}
	}
	
	for _, key := range keysToDelete {
		delete(store, key)
	}
	
	if len(keysToDelete) > 0 {
		log.Printf("Cleaned up %d expired keys", len(keysToDelete))
	}
}

// StartCleanupRoutine starts a background goroutine that periodically cleans up expired keys
func StartCleanupRoutine() {
	go func() {
		ticker := time.NewTicker(1 * time.Second) // Check every second
		defer ticker.Stop()
		
		for range ticker.C {
			CleanupExpiredKeys()
		}
	}()
	log.Println("Started background cleanup routine for expired keys")
}