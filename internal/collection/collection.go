package collection

import (
	"time"

	"github.com/YarikRevich/HideSeek-Server/internal/api"
	// "github.com/google/uuid"
)

var instance *Collection

type Game struct {
	Started bool
}

type PC struct {
	Data *api.PC
	AddTime time.Time
}

type Collection struct {
	Games map[string]Game
	Worlds map[string]*api.World
	PCs map[string]map[string]*PC
	Elements map[string]map[string]*api.Element
	Weapons map[string]map[string]*api.Weapon
	Ammo map[string]map[string]*api.Ammo
}

func (c *Collection) CleanDataByUUID(u string){
	delete(c.Worlds, u)
	delete(c.PCs, u)
	delete(c.Elements, u)
	delete(c.Weapons, u)
	delete(c.Ammo, u)
}

func UseCollection() *Collection{
	if instance == nil{
		instance = &Collection{
			Games: make(map[string]Game),
			Worlds: make(map[string]*api.World),
			PCs: make(map[string]map[string]*PC),
			Elements: make(map[string]map[string]*api.Element),
			Weapons: make(map[string]map[string]*api.Weapon),
			Ammo: make(map[string]map[string]*api.Ammo),
		}
	}
	return instance
}