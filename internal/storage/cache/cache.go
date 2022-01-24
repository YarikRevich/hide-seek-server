package cache

import (
	"time"

	"github.com/YarikRevich/hide-seek-server/tools/params"
)

//Caching works as Pub/Sub service
type Cache struct {
	mainTicker, stubTicker *time.Ticker
	storage                []func()
}

func (c *Cache) Subscribe(b func()) {
	c.storage = append(c.storage, b)
}

func (c *Cache) Run() {
	go func() {
		for {
			select {
			case <-c.mainTicker.C:
				for _, v := range c.storage {
					v()
				}
			case <-c.stubTicker.C:
			}
		}
	}()
}

func New() *Cache {
	return &Cache{
		mainTicker: time.NewTicker(time.Duration(params.GetCacheTime()) * time.Millisecond),
		stubTicker: time.NewTicker(500 * time.Millisecond),
	}
}
