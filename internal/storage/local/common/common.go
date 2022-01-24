package common

import "time"

type Collection interface {
	InsertOrUpdate(string, interface{})

	//Key param can be represented as a single
	//key or a key sequence if value of the storage map
	//is an array and you should delete or find a proper value
	Find(key interface{}) interface{}
	Delete(key interface{})

	//Returns callback which will be used by
	//caching service
	//If collection does not use caching it can
	//be left empty
	Cache() func()
}

type CollectionEntity struct {
	Data      interface{}
	Timestamp time.Time
}
