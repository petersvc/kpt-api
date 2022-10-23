package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetGpuCollection() *mongo.Collection {
	// fmt.Println("getting the Gpus Collection on mongoDB!")
	Db, collectionName := DbService()
	// fmt.Println("collection retrieved!")
	return Db.Collection(collectionName[0])
}

func GetGpuFilterCollection() *mongo.Collection {
	// fmt.Println("getting the Filters Collection on mongoDB!")
	Db, collectionName := DbService()
	// fmt.Println("collection retrieved!")
	return Db.Collection(collectionName[1])
}

func GetCpuCollection() *mongo.Collection {
	// fmt.Println("getting the Gpus Collection on mongoDB!")
	Db, collectionName := DbService()
	// fmt.Println("collection retrieved!")
	return Db.Collection(collectionName[2])
}
