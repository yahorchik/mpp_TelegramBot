package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

type Message struct {
	Data int
	Text string
	User int64
}

func InitCache() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	Cache = c
}

/*func (lc Cache) InitCache() *cache.Cache {
	lc.c = cache.New(5*time.Minute, 10*time.Minute)
	return lc.c
}*/
