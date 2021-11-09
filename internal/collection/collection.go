package collection

import "github.com/google/uuid"

var instance *Collection

type Collection struct {
	worlds map[uuid.UUID]interface{}
	pcs map[uuid.UUID]interface{}
	elements map[uuid.UUID]interface{}
	weapons map[uuid.UUID]interface{}
	ammo map[uuid.UUID]interface{}
}

// func 

// func (c *Collection) Delete()

func (c *Collection) CleanDataByUUID(u uuid.UUID){
	delete(c.worlds, u)
	delete(c.pcs, u)
	delete(c.elements, u)
	delete(c.weapons, u)
	delete(c.ammo, u)
}

func UseCollection() *Collection{
	if instance == nil{
		instance = &Collection{
			worlds: make(map[uuid.UUID]interface{}),
			pcs: make(map[uuid.UUID]interface{}),
			elements: make(map[uuid.UUID]interface{}),
			weapons: make(map[uuid.UUID]interface{}),
			ammo: make(map[uuid.UUID]interface{}),
		}
	}
	return instance
}