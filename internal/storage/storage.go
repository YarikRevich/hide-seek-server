package storage

import (
	"github.com/YarikRevich/HideSeek-Server/internal/storage/cache"
	"github.com/YarikRevich/HideSeek-Server/internal/storage/db"
)

var instance *Storage

type Storage struct {
	db    *db.DB
	cache *cache.Cache
}

func (s *Storage) DB() *db.DB {
	return s.db
}

func (s *Storage) Cache() *cache.Cache {
	return s.cache
}

func UseStorage() *Storage {
	if instance == nil {
		instance = &Storage{
			db:    db.New(),
			cache: cache.New(),
		}
	}
	return instance
}
