package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	content map[string]CacheEntry
	m       *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	defer c.m.Unlock()
	c.content[key] = CacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	content, ok := c.content[key]
	if !ok {
		return nil, ok
	}
	return content.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.m.Lock()
	defer c.m.Unlock()
	for k, v := range c.content {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.content, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{content: make(map[string]CacheEntry), m: &sync.Mutex{}}
	go cache.reapLoop(interval)
	return cache
}
