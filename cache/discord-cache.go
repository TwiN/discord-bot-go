package cache

import "log"

func init() {
	log.Println("[discord-cache][init] Initializing discord cache in global cache")
	cache["channel"] = make(map[string][]string)
	cache["guild"] = make(map[string][]string)
}
