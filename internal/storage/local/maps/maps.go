package maps

import (
	"github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/HideSeek-Server/internal/storage/local/common"
)

type MapsCollection struct {
	elements map[string]*proto.Map
}

func (mc *MapsCollection) InsertOrUpdate(key string, data interface{}) {
	mc.elements[key] = data.(*proto.Map)
}

func (mc *MapsCollection) Find(key string) interface{} {
	v, ok := mc.elements[key]
	if ok {
		return v
	}
	return []*proto.Map{}
}

func (mc *MapsCollection) Delete(key string) {
	delete(mc.elements, key)
}

func New() common.Collection {
	return &MapsCollection{
		elements: make(map[string]*proto.Map),
	}
}
