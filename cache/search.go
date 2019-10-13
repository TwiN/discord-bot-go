package cache

import "log"

var (
	Google  *Cache
	Youtube *Cache
	Urban   *Cache
)

func init() {
	log.Println("[cache][search][init] Initializing search cache in global cache")
	Google = NewCache()
	Youtube = NewCache()
	Urban = NewCache()
}
