package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex
	lifetime time.Duration
}

type cacheEntry struct {
	value []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache: map[string]cacheEntry{},
		mu: sync.Mutex{},
		lifetime: interval,
	}
	go c.reapLoop(c.lifetime)
	return &c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		value: value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if val, ok := c.cache[key]; ok {
		return val.value, ok
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	for k := range c.cache {
		if time.Since(c.cache[k].createdAt) >= interval {
			delete(c.cache, k)
		}
	}
	c.mu.Unlock()
}