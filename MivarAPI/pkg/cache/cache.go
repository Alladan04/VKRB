package cache

import (
	"sync"
)

// Cache represents an in-memory key-value store.
type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// NewCache creates and initializes a new Cache instance.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set adds or updates a key-value pair in the cache.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get retrieves the value associated with the given key from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

// Delete removes a key-value pair from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Clear removes all key-value pairs from the cache.
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]interface{})
}

//func main() {
//	cache := NewCache()
//
//	// Adding data to the cache
//	cache.Set("key1", "value1")
//	cache.Set("key2", 123)
//
//	// Retrieving data from the cache
//	if val, ok := cache.Get("key1"); ok {
//		fmt.Println("Value for key1:", val)
//	}
//
//	// Deleting data from the cache
//	cache.Delete("key2")
//
//	// Clearing the cache
//	cache.Clear()
//
//	time.Sleep(time.Second) // Sleep to allow cache operations to complete
//}
