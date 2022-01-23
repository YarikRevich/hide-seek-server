package cache

// var ticker = time.NewTicker(time.Duration(params.GetCacheTime()))

type Cache struct {
}

// func (c *Cache) start() {
// 	go func() {
// 		for range c.ticker.C {
// 			for k, v := range c.cache {
// 				if time.Since(v) == 0 {
// 					// collection.UseCollection().CleanDataByUUID(k)
// 					delete(c.cache, k)
// 				}
// 			}
// 		}
// 	}()
// }

// func (c *Cache) Postpone(u string) {
// 	c.cache[u] = time.Now().Add(time.Minute * 5)
// }

func New() *Cache {
	return new(Cache)
}
