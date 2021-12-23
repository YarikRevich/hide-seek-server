package elements

import (
	"github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/HideSeek-Server/internal/storage/local/common"
)

type ElementsCollection struct {
	elements map[string][]*proto.Element
}

func (ec *ElementsCollection) InsertOrUpdate(key string, data interface{}) {
	ec.elements[key] = append(ec.elements[key], data.(*proto.Element))
}

func (ec *ElementsCollection) Find(key string) interface{} {
	v, ok := ec.elements[key]
	if ok {
		return v
	}
	return []*proto.Element{}
}

func (ec *ElementsCollection) Delete(key string) {
	delete(ec.elements, key)
}

func New() common.Collection {
	return &ElementsCollection{
		elements: make(map[string][]*proto.Element),
	}
}
