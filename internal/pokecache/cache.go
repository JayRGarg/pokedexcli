package pokecache

import (
    "time"
    "sync"
)


type Cache struct {
    interval        time.Duration
    UrlToEntry      map[string]cacheEntry
    stop            chan bool
    mu              sync.Mutex
}

type cacheEntry struct {
    createdAt       time.Time
    val             []byte
}

func NewCache(interval time.Duration) *Cache {
    c := Cache {
        interval: interval,
        UrlToEntry: make(map[string]cacheEntry),
        stop: make(chan bool),
    }
    go c.reapLoop()
    return &c
}

func (c *Cache) Stop() {
    close(c.stop)
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.UrlToEntry[key] = cacheEntry{
        createdAt: time.Now(),
        val: val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    entry, exists := c.UrlToEntry[key]
    if !exists {
        return nil, exists
    }
    return entry.val, exists
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            c.mu.Lock()
            for k, v := range c.UrlToEntry {
                if time.Since(v.createdAt) > c.interval {
                    delete(c.UrlToEntry, k)
                }
            }
            c.mu.Unlock()
        case <-c.stop:
            return
        }
    }
}
