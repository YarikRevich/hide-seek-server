package storage

import (
	"github.com/YarikRevich/HideSeek-Server/internal/storage/external-api"
	"github.com/YarikRevich/HideSeek-Server/internal/storage/internal-api"
	// "github.com/google/uuid"
)

// var instance *Collection

// type Game struct {
// 	Started bool
// }

// type PC struct {
// 	Data *api.PC
// 	AddTime time.Time
// }

// type Collection struct {
// 	Games map[string]Game
// 	Worlds map[string]*api.World
// 	PCs map[string]map[string]*PC
// 	Elements map[string]map[string]*api.Element
// 	Weapons map[string]map[string]*api.Weapon
// 	Ammo map[string]map[string]*api.Ammo
// }

// func (c *Collection) CleanDataByUUID(u string){
// 	delete(c.Worlds, u)
// 	delete(c.PCs, u)
// 	delete(c.Elements, u)
// 	delete(c.Weapons, u)
// 	delete(c.Ammo, u)
// }

var instance *Storage

type Storage struct {
	*externalapistorage.ExternalApiStorage
	*internalapistorage.InternalApiStorage
}

func UseStorage() *Storage {
	if instance == nil {
		instance = &Storage{
			externalapistorage.NewExternalApiStorage(),
			internalapistorage.NewInternalApiStorage()}
	}
	return instance
}
