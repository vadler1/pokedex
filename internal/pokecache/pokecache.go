package pokecache

type Cache struct {
	data map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.data[key] = value
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, ok := c.data[key]

	if !ok {
		return nil, false
	}

	return value, ok
}
