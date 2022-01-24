package pcs

import (
	"github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto"
	"github.com/YarikRevich/hide-seek-server/internal/monitoring"
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
)

type PCsCollection struct {
	elements map[string][]*proto.PC
}

func (mc *PCsCollection) InsertOrUpdate(key string, data interface{}) {
	q := data.(*proto.PC)
	for i, v := range mc.elements[key] {
		if v.Base.GetId() == q.Base.GetId() {
			mc.elements[key] = append(mc.elements[key][:i], mc.elements[key][i+1:]...)
			mc.elements[key] = append(mc.elements[key], data.(*proto.PC))
			break
		}
	}
	monitoring.UseMonitoring().RegisterManager().PlayersGauge().Inc()
	q.LobbyNumber = int64(len(mc.elements[key]) + 1)
	mc.elements[key] = append(mc.elements[key], data.(*proto.PC))
}

func (mc *PCsCollection) Find(key interface{}) interface{} {
	if v, ok := mc.elements[key.(string)]; ok {
		return v
	}
	return []*proto.PC{}
}

func (mc *PCsCollection) Delete(key interface{}) {
	delete(mc.elements, key.(string))
	monitoring.UseMonitoring().RegisterManager().PlayersGauge().Dec()
}

func (mc *PCsCollection) Cache() func() {
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
	return &PCsCollection{
		elements: make(map[string][]*proto.PC),
	}
}
