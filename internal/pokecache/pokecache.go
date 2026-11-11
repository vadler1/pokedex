package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

type CacheData map[string]CacheEntry

type Cache struct {
	data CacheData
	mu   *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: make(CacheData),
		mu:   &sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
