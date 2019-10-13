package cache

type Cache struct {
	cache map[string][]string
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string][]string)}
}

func (c *Cache) Get(key string) []string {
	return c.cache[key]
}

func (c *Cache) Put(key string, value []string) {
	c.cache[key] = value
}

func (c *Cache) Has(key string) bool {
	return c.cache[key] != nil
}
