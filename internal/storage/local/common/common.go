package common

type Collection interface {
	InsertOrUpdate(string, interface{})
	Find(string) interface{}
	Delete(string)
}
