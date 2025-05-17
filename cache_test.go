package main

import (
    "testing"
    "fmt"
    "time"
    "sync"
	"github.com/jayrgarg/pokedexcli/internal/pokecache"
)


func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
            t.Cleanup(func() { cache.Stop() })
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
    t.Cleanup(func() { cache.Stop() })
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestCacheMiss(t *testing.T) {
    cache := pokecache.NewCache(5 * time.Second)
    _, ok := cache.Get("nonexistent-key")
    if ok {
        t.Error("Expected cache miss for nonexistent key, but got a hit")
    }
}

func TestOverwriteEntry(t *testing.T) {
    cache := pokecache.NewCache(5 * time.Second)
    key := "test-key"
    
    cache.Add(key, []byte("original-value"))
    cache.Add(key, []byte("new-value"))
    
    val, ok := cache.Get(key)
    if !ok {
        t.Error("Expected to find key after overwrite")
        return
    }
    
    if string(val) != "new-value" {
        t.Errorf("Expected overwritten value 'new-value', but got '%s'", string(val))
    }
}

func TestConcurrentAccess(t *testing.T) {
    cache := pokecache.NewCache(5 * time.Second)
    var wg sync.WaitGroup
    
    // Launch multiple goroutines
    for i := range 100 {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            key := fmt.Sprintf("key-%d", id)
            cache.Add(key, []byte(fmt.Sprintf("value-%d", id)))
            
            // Read our own key
            val, ok := cache.Get(key)
            if !ok || string(val) != fmt.Sprintf("value-%d", id) {
                t.Errorf("Concurrent access failure for key %s", key)
            }
        }(i)
    }
    
    wg.Wait()
}

func TestEdgeCases(t *testing.T) {
    cache := pokecache.NewCache(5 * time.Second)
    
    // Test empty key
    cache.Add("", []byte("empty-key-value"))
    val, ok := cache.Get("")
    if !ok || string(val) != "empty-key-value" {
        t.Error("Failed with empty key")
    }
    
    // Test large value
    largeVal := make([]byte, 1024*1024) // 1MB
    cache.Add("large-key", largeVal)
    val, ok = cache.Get("large-key")
    if !ok || len(val) != 1024*1024 {
        t.Error("Failed with large value")
    }
}

func TestIntervalVariations(t *testing.T) {
    // Test with very short interval
    t.Run("Very short interval", func(t *testing.T) {
        cache := pokecache.NewCache(10 * time.Millisecond)
        cache.Add("key", []byte("value"))
        
        time.Sleep(15 * time.Millisecond)
        
        _, ok := cache.Get("key")
        if ok {
            t.Error("Expected entry to be reaped with short interval")
        }
    })
    
    // Test with longer interval
    t.Run("Longer interval", func(t *testing.T) {
        cache := pokecache.NewCache(100 * time.Millisecond)
        cache.Add("key", []byte("value"))
        
        time.Sleep(50 * time.Millisecond) // Half the interval
        
        _, ok := cache.Get("key")
        if !ok {
            t.Error("Expected entry to still exist before reap interval")
        }
    })
}

func TestMultipleEntriesReaping(t *testing.T) {
    const interval = 50 * time.Millisecond
    cache := pokecache.NewCache(interval)
    
    // Add multiple entries
    for i := range 5 {
        key := fmt.Sprintf("key-%d", i)
        cache.Add(key, []byte(fmt.Sprintf("value-%d", i)))
    }
    
    // Verify all entries exist
    for i := range 5 {
        key := fmt.Sprintf("key-%d", i)
        _, ok := cache.Get(key)
        if !ok {
            t.Errorf("Expected to find key %s before reaping", key)
        }
    }
    
    // Wait for the entries to be reaped
    time.Sleep(interval + 10*time.Millisecond)
    
    // Verify all entries have been reaped
    for i := range 5 {
        key := fmt.Sprintf("key-%d", i)
        _, ok := cache.Get(key)
        if ok {
            t.Errorf("Expected key %s to be reaped", key)
        }
    }
}
