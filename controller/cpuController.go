package controller

import (
	"context"
	"encoding/json"
	"fmt"
	service "kpt_api/services"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCpuData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	collection := service.GetCpuCollection()
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}

	var gpus []primitive.M
	for cur.Next(context.Background()) {
		var gpu primitive.M
		err := cur.Decode(&gpu)
		if err != nil {
			panic(err)
		}
		gpus = append(gpus, gpu)
	}
	fmt.Println("All Gpus retrieved!")
	defer cur.Close(context.Background())
	json.NewEncoder(w).Encode(gpus)
	defer r.Body.Close()
	if r.Close {
		fmt.Println("request closed")
	}
}
