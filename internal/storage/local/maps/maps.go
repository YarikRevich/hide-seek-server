package maps

import (
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
)

type MapsCollection struct {
	elements map[string]*proto.Map
}

func (mc *MapsCollection) InsertOrUpdate(key string, data interface{}) {
	mc.elements[key] = data.(*proto.Map)
}

func (mc *MapsCollection) Find(key interface{}) interface{} {
	if v, ok := mc.elements[key.(string)]; ok {
		return v
	}
	return []*proto.Map{}
}

func (mc *MapsCollection) Delete(key interface{}) {
	delete(mc.elements, key.(string))
}

func New() common.Collection {
	return &MapsCollection{
		elements: make(map[string]*proto.Map),
	}
}
