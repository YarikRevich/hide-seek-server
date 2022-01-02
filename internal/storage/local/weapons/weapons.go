package weapons

import (
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
)

type WeaponsCollection struct {
	elements map[string][]*proto.Weapon
}

func (mc *WeaponsCollection) InsertOrUpdate(key string, data interface{}) {

	mc.elements[key] = append(mc.elements[key], data.(*proto.Weapon))
}

func (mc *WeaponsCollection) Find(key string) interface{} {
	v, ok := mc.elements[key]
	if ok {
		return v
	}
	return []*proto.Weapon{}
}

func (mc *WeaponsCollection) Delete(key string) {
	delete(mc.elements, key)
}

func New() common.Collection {
	return &WeaponsCollection{
		elements: make(map[string][]*proto.Weapon),
	}
}
