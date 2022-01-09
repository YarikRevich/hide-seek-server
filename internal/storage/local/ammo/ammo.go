package ammo

import (
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
)

type AmmoCollection struct {
	ammo map[string][]*proto.Ammo
}

func (ac *AmmoCollection) InsertOrUpdate(key string, data interface{}) {
	ac.ammo[key] = append(ac.ammo[key], data.(*proto.Ammo))
}

func (ac *AmmoCollection) Find(key interface{}) interface{} {
	if v, ok := ac.ammo[key.(string)]; ok {
		return v
	}
	return []*proto.Ammo{}
}

func (ac *AmmoCollection) Delete(key interface{}) {
	delete(ac.ammo, key.(string))
}

func New() common.Collection {
	return &AmmoCollection{
		ammo: make(map[string][]*proto.Ammo),
	}
}
