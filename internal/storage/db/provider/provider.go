package provider

import (
	"context"
	"fmt"

	"github.com/YarikRevich/hide-seek-server/tools/params"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Provider struct {
	client *mongo.Client
}

func (p *Provider) Connect() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", params.GetMongoHost())))
	if err != nil {
		logrus.Fatal(err)
	}
	p.client = client
}

func (p *Provider) Collection(collection string) *mongo.Collection {
	return p.client.Database(params.GetMongoDB()).Collection(collection)
}

func New() *Provider {
	return new(Provider)
}
