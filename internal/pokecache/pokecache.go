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
	data         CacheData
	mu           *sync.RWMutex
	ticker       *time.Ticker
	reapInterval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)

	cache := &Cache{
		data:         make(CacheData),
		mu:           &sync.RWMutex{},
		ticker:       ticker,
		reapInterval: interval,
	}

	go func() {
		cache.reapLoop()
	}()

	return cache
}

func (c *Cache) reapLoop() {
	for range c.ticker.C {
		c.mu.Lock()
		for k, v := range c.data {
			if time.Since(v.createdAt) > c.reapInterval {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Cleanup() {
	c.ticker.Stop()
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
