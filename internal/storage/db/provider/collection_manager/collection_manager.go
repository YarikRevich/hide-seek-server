package collectionmanager

import (
	"context"
	"reflect"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionManager struct {
	collection *mongo.Collection
}

func (cm *CollectionManager) InsertOne(data bson.M) {
	if _, err := cm.collection.InsertOne(context.Background(), data); err != nil {
		logrus.Fatal(err)
	}
}

func (cm *CollectionManager) FindOne(filter bson.D, dst interface{}) {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr {
		logrus.Fatal("dst in findone must be ptr")
	}
	if err := cm.collection.FindOne(context.Background(), filter).Decode(dst); err != nil {
		logrus.Fatal(err)
	}
}

func New(collection *mongo.Collection) *CollectionManager {
	return &CollectionManager{collection}
}
