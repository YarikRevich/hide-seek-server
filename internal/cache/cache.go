package cache

import (
	"time"
	// "github.com/YarikRevich/hide-seek-server/internal/collection"
)

var instance *Cache

type Cache struct {
	ticker *time.Ticker
	cache  map[string]time.Time
}

func (c *Cache) start() {
	go func() {
		spain := time.NewTicker(time.Millisecond * 10)
		for {
			select {
			case <-c.ticker.C:
				for k, v := range c.cache {
					if time.Since(v) == 0 {
						// collection.UseCollection().CleanDataByUUID(k)
						delete(c.cache, k)
					}
				}
			case <-spain.C:
			}
		}
	}()
}

func (c *Cache) Postpone(u string) {
	c.cache[u] = time.Now().Add(time.Minute * 5)
}

func UseCache() *Cache {
	if instance == nil {
		instance = &Cache{
			ticker: time.NewTicker(time.Second),
			cache:  make(map[string]time.Time),
		}
		instance.start()
	}
	return instance
}
