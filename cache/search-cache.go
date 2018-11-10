package cache

import "log"

var cache = make(map[string]map[string][]string)

func init() {
	log.Println("[search-cache][init] Initializing global cache")
	cache["google"] = make(map[string][]string)
	cache["youtube"] = make(map[string][]string)
	cache["urban"] = make(map[string][]string)
}

func Get(cacheName string, key string) []string {
	return cache[cacheName][key]
}

func Put(cacheName string, key string, value []string) {
	cache[cacheName][key] = value
}

func Has(cacheName string, key string) bool {
	return cache[cacheName][key] != nil
}