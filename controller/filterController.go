package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"kpt_api/services"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFilterData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// // Validations
	// services.ValidateQuery(r.RequestURI)

	// Filter and the query options
	filter := bson.D{}
	opts := options.Find()

	// // Get the limit of items and the page
	// limit := services.GetLimit(r)
	// page := services.GetPage(r)

	// // SORT
	// sortkey, sortType := services.GetSortOptions(r)
	// if sortkey != "" {
	// 	opts.SetSort(bson.D{{Key: sortkey, Value: sortType}})
	// }

	// // LIMIT
	// opts.SetLimit(limit)

	// // PAGINATION
	// if page > 0 {
	// 	opts.SetSkip((page - 1) * limit)
	// }

	// Get the Collection
	collection := services.GetGpuFilterCollection()

	// Find the data, run the Query
	cur, err := collection.Find(context.TODO(), filter, opts)

	if err != nil {
		panic(err)
	}

	var data []primitive.M
	for cur.Next(context.Background()) {
		var element primitive.M
		err := cur.Decode(&element)
		if err != nil {
			panic(err)
		}
		data = append(data, element)
	}

	if len(data) < 1 {
		fmt.Println("No data found")
	}

	fmt.Println()

	defer cur.Close(context.Background())
	defer r.Body.Close()
	json.NewEncoder(w).Encode(data)
}
