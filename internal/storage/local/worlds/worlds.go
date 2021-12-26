package worlds

import (
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
)

type WorldsCollection struct {
	elements map[string][]*proto.World
}

func (mc *WorldsCollection) InsertOrUpdate(key string, data interface{}) {
	mc.elements[key] = append(mc.elements[key], data.(*proto.World))
}

func (mc *WorldsCollection) Find(key string) interface{} {
	v, ok := mc.elements[key]
	if ok {
		return v
	}
	return []*proto.World{}
}

func (mc *WorldsCollection) Delete(key string) {
	delete(mc.elements, key)
}

func New() common.Collection {
	return &WorldsCollection{
		elements: make(map[string][]*proto.World),
	}
}
