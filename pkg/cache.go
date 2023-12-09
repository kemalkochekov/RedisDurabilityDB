package pkg

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	store map[string]*cacheItem
	mu    sync.RWMutex
}
type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]*cacheItem),
	}
}
func (c *Cache) Get(ctx context.Context, key string) (any, error) {
	c.mu.RLock()
	item, found := c.store[key]
	c.mu.RUnlock()
	if !found {
		return "", fmt.Errorf("Not Found")
	}
	// Check if the item has expired
	expired := found && time.Now().After(item.expiration)
	if expired {
		c.mu.Lock()
		delete(c.store, key)
		c.mu.Unlock()
	}
	return item.value, nil
}

func (c *Cache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = &cacheItem{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
	return nil
}
