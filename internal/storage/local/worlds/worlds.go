package worlds

import (
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
)

type WorldsCollection struct {
	elements map[string]*proto.World
}

func (mc *WorldsCollection) InsertOrUpdate(key string, data interface{}) {
	mc.elements[key] = data.(*proto.World)
}

func (mc *WorldsCollection) Find(key interface{}) interface{} {
	if v, ok := mc.elements[key.(string)]; ok {
		return v
	}
	return []*proto.World{}
}

func (mc *WorldsCollection) Delete(key interface{}) {
	delete(mc.elements, key.(string))
}

func New() common.Collection {
	return &WorldsCollection{
		elements: make(map[string]*proto.World),
	}
}
