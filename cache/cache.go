package cache

var cache = make(map[string]map[string][]string)


func Get(cacheName string, key string) []string {
	return cache[cacheName][key]
}


func Put(cacheName string, key string, value []string) {
	cache[cacheName][key] = value
}


func Has(cacheName string, key string) bool {
	return cache[cacheName][key] != nil
}
