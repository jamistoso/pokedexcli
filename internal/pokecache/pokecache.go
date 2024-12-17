package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cacheMux		*sync.Mutex
	cacheMap		map[string]cacheEntry
	interval		time.Duration
}

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

func NewCache(duration time.Duration) (Cache) {
	cache := Cache {
		cacheMux:	&sync.Mutex{},
		cacheMap:	map[string]cacheEntry{},
		interval:	duration,
	}
	go cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: 	time.Now(),
		val:		val,
	}
	c.cacheMux.Lock()
	c.cacheMap[key] = entry
	c.cacheMux.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.cacheMux.Lock()
	defer c.cacheMux.Unlock()
	entry, exists := c.cacheMap[key]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.cacheMux.Lock()
		for key := range c.cacheMap {
			if time.Since(c.cacheMap[key].createdAt) > c.interval {
				delete(c.cacheMap, key)
			}
		}
		c.cacheMux.Unlock()
	}
}