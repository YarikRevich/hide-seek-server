package storage

import (
	"github.com/YarikRevich/hide-seek-server/internal/storage/cache"
	"github.com/YarikRevich/hide-seek-server/internal/storage/db"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local"
)

var instance *Storage

type Storage struct {
	db    *db.DB
	cache *cache.Cache
	local *local.Local
}

func (s *Storage) DB() *db.DB {
	return s.db
}

func (s *Storage) Cache() *cache.Cache {
	return s.cache
}

func (s *Storage) Local() *local.Local {
	return s.local
}

func UseStorage() *Storage {
	if instance == nil {
		c := cache.New()
		l := local.New(c)
		instance = &Storage{
			db:    db.New(),
			cache: c,
			local: l,
		}

	}
	return instance
}
