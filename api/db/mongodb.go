package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

func GetClient(dbUri string) *mongo.Client {
	var clientInstance *mongo.Client
	var mongoOnce sync.Once

	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(dbUri)
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Ping(context.Background(), nil)
		if err != nil {
			fmt.Println("Unable to connect to MongoDB.  Please check your settings..")
			panic(err)
		}
		clientInstance = client
	})
	return clientInstance
}

func GetDatabase(dbUri string, dbName string) *mongo.Database {
	client := GetClient(dbUri)
	return client.Database(dbName)
}
