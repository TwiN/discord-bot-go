package cache

import (
	"log"
)

var (
	Channel *Cache
	Guild   *Cache
)

func init() {
	log.Println("[cache][discord][init] Initializing discord cache in global cache")
	Channel = NewCache()
	Guild = NewCache()
}
