package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	Value      string
	Expiration int64
}

type Cache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

func (c *Cache) Set(key, value string, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration).Unix(),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found || time.Now().Unix() > item.Expiration {
		return "", false
	}
	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *Cache) Cleaner() {
	for {
		time.Sleep(5 * time.Second)
		c.mu.Lock()
		for k, v := range c.items {
			if time.Now().Unix() > v.Expiration {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}

func main() {
	cache := Cache{items: make(map[string]CacheItem)}
	go cache.Cleaner()

	fmt.Println("⚡ Simple Cache Server (with TTL)")
	fmt.Println("--------------------------------")

	for {
		fmt.Println("\nCommands: set | get | del | exit")
		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "set":
			var key, value string
			var ttl int
			fmt.Print("Key: ")
			fmt.Scan(&key)
			fmt.Print("Value: ")
			fmt.Scan(&value)
			fmt.Print("TTL (seconds): ")
			fmt.Scan(&ttl)
			cache.Set(key, value, time.Duration(ttl)*time.Second)
			fmt.Println("✅ Cached!")

		case "get":
			var key string
			fmt.Print("Key: ")
			fmt.Scan(&key)
			if val, ok := cache.Get(key); ok {
				fmt.Println("🔍 Value:", val)
			} else {
				fmt.Println("❌ Key expired or not found.")
			}

		case "del":
			var key string
			fmt.Print("Key: ")
			fmt.Scan(&key)
			cache.Delete(key)
			fmt.Println("🗑️ Deleted.")

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Unknown command.")
		}
	}
}
