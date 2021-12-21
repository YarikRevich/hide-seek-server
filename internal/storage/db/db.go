package db

import (
	"github.com/YarikRevich/HideSeek-Server/internal/storage/db/provider"
	collectionmanager "github.com/YarikRevich/HideSeek-Server/internal/storage/db/provider/collection_manager"
)

type DB struct {
	provider *provider.Provider
}

func (d *DB) History() *collectionmanager.CollectionManager {
	return collectionmanager.New(d.provider.Collection("history"))
}

func (d *DB) Profiles() *collectionmanager.CollectionManager {
	return collectionmanager.New(d.provider.Collection("profiles"))
}

func New() *DB {
	return &DB{
		provider: provider.New(),
	}
}
