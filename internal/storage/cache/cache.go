package cache

type Cache struct{}

func New() *Cache {
	return new(Cache)
}
