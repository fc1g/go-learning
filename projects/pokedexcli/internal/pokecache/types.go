package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]CacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
