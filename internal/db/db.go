package db

var instance *DB

type DB struct {}

func UseDB()*DB{
	if instance == nil{
		instance = new(DB)
	}
	return instance
}