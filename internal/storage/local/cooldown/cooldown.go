package cooldown

import (
	"github.com/YarikRevich/hide-seek-server/internal/storage/local/common"
	"github.com/sirupsen/logrus"
)

type CooldownCollection struct {
	elements map[string][]string
}

func (mc *CooldownCollection) InsertOrUpdate(key string, data interface{}) {
	mc.elements[key] = append(mc.elements[key], data.(string))
}

//Find method for cooldown allows only double array or a single string
//Returns boolean type(whether value exists by this key or not)
func (mc *CooldownCollection) Find(key interface{}) interface{} {
	switch x := key.(type) {
	case string:
		_, ok := mc.elements[x]
		return ok
	case [2]string:
		for _, v := range mc.elements[x[0]] {
			if v == x[1] {
				return true
			}
		}
	}
	return false
}

//Delete method for cooldown allows only double array or a single string
func (mc *CooldownCollection) Delete(key interface{}) {
	switch x := key.(type) {
	case string:
		delete(mc.elements, x)
	case [2]string:
		for i, v := range mc.elements[x[0]] {
			if v == x[1] {
				mc.elements[x[0]] = append(mc.elements[x[0]][:i], mc.elements[x[0]][i+1:]...)
			}
		}
	default:
		logrus.Fatal("allowed only double array maximally or single one at least")
	}
}

func (mc *CooldownCollection) Cache() func() {
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
	return &CooldownCollection{
		elements: make(map[string][]string),
	}
}
