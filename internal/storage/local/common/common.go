package common

type Collection interface {
	InsertOrUpdate(string, interface{})

	//Key param can be represented as a single
	//key or a key sequence if value of the storage map
	//is an array and you should delete or find a proper value
	Find(key interface{}) interface{}
	Delete(key interface{})
}
