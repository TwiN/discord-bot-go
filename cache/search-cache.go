package cache

import "log"


func init() {
	log.Println("[search-cache][init] Initializing search cache in global cache")
	cache["google"] = make(map[string][]string)
	cache["youtube"] = make(map[string][]string)
	cache["urban"] = make(map[string][]string)
}
