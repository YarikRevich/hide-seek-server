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

func (ac *AmmoCollection) Find(key string) interface{} {
	v, ok := ac.ammo[key]
	if ok {
		return v
	}
	return []*proto.Ammo{}
}

func (ac *AmmoCollection) Delete(key string) {
	delete(ac.ammo, key)
}

func New() common.Collection {
	return &AmmoCollection{
		ammo: make(map[string][]*proto.Ammo),
	}
}
