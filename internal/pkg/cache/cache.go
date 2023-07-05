package cache

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

func MsgToCache(string) {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Items()
	for dick, item := range c.Items() {
		log.Println(dick, item)
		switch item.Object.(type) {
		case int:

		}
	}
}
