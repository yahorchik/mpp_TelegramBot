package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

//var Cache *cache.Cache
//
//type Message struct {
//	Data int
//	Text string
//	User int64
//}

type Cache struct {
	C *cache.Cache
}

const (
	defaultExpirationInterval = 5 * time.Minute
	defaultCleanupInterval    = 10 * time.Minute
)

func InitCache() *Cache {
	return &Cache{C: cache.New(defaultExpirationInterval, defaultCleanupInterval)}
}
