package services

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbService() (*mongo.Database, []string) {
	// fmt.Println("DbService is getting started!")

	DbSettings := GetDbSettings()
	ClientOption := options.Client().ApplyURI(DbSettings.ConnectUri)
	client, err := mongo.Connect(context.TODO(), ClientOption)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Connected to DB!")

	return client.Database(DbSettings.DatabaseName), DbSettings.CollectionName
}
