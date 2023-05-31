package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (cache *Cache) add(key string, value []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()

	cache.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (cache *Cache) get(key string) ([]byte, bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()

	c, ok := cache.cache[key]
	return c.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	for range tick.C {
		t := time.Now().Add(-interval)
		fmt.Println(t)
		for i, c := range cache.cache {
			if c.createdAt.Before(t) {
				delete(cache.cache, i)
			}
		}
	}
}
