package cfg

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "mydb"
const mongoUri = "mongodb://root:example@localhost:27017/"

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MongoDb MongoInstance

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	db := client.Database(dbName)

	MongoDb = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func Disconnect() error {
	return MongoDb.Client.Disconnect(context.Background())
}
