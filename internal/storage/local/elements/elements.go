package elements

import (
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
)

type ElementsCollection struct {
	elements map[string][]*proto.Element
}

func (ec *ElementsCollection) InsertOrUpdate(key string, data interface{}) {
	ec.elements[key] = append(ec.elements[key], data.(*proto.Element))
}

func (ec *ElementsCollection) Find(key interface{}) interface{} {
	v, ok := ec.elements[key.(string)]
	if ok {
		return v
	}
	return []*proto.Element{}
}

func (ec *ElementsCollection) Delete(key interface{}) {
	delete(ec.elements, key.(string))
}

func (ec *ElementsCollection) Cache() func() {
	return func() {
		// for k, v := range c.cache {
		// if time.Since(v) == 0 {
		// collection.UseCollection().CleanDataByUUID(k)
		// delete(c.cache, k)
		// }
		// }
	}
}

func New() common.Collection {
	return &ElementsCollection{
		elements: make(map[string][]*proto.Element),
	}
}
