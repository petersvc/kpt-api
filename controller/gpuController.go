package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"kpt_api/model"
	"kpt_api/services"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetGpuData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Validations
	services.ValidateQuery(r.RequestURI)

	// Filter and the query options

	// Set the Options
	opts := options.Find()

	// Get the limit of items and the page
	limit := services.GetLimit(r)
	page := services.GetPage(r)

	// SORT
	sortkey, sortType := services.GetSortOptions(r)
	if sortkey != "" {
		opts.SetSort(bson.D{{Key: sortkey, Value: sortType}})
	}

	// LIMIT
	opts.SetLimit(limit)

	// PAGINATION
	if page > 0 {
		opts.SetSkip((page - 1) * limit)
	}

	// Get the Collection
	collection := services.GetGpuCollection()

	// Find the data, filter and run the Query
	var cur *mongo.Cursor

	if strings.Contains(r.RequestURI, "&") {
		parameters := services.GetFilters(r)

		filter := bson.M{
			"$and": []bson.M{},
		}

		fmt.Println("filter:", filter)

		for _, v := range parameters {
			var filters []bson.M
			parameter := v
			fmt.Println("parameter:", parameter.Key)
			if len(parameter.Value) > 1 {
				for _, v := range parameter.Value {
					fmt.Println("value:", v)
					filter1 := bson.M{parameter.Key: v}
					filters = append(filters, filter1)
				}
			} else {
				fmt.Println("value:", parameter.Value[0])
				filter1 := bson.M{parameter.Key: parameter.Value[0]}
				filters = append(filters, filter1)
			}
			filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"$or": filters})
			fmt.Println("filters:", filters)
		}

		fmt.Println("filter:", filter)

		var err error
		cur, err = collection.Find(context.TODO(), filter, opts)

		services.GetError(err)

	} else {
		fmt.Println("else:", strings.Contains(r.RequestURI, "&"))
		var err error
		cur, err = collection.Find(context.TODO(), bson.D{}, opts)
		services.GetError(err)
	}

	// cur, err := collection.Find(context.TODO(), filter, opts)

	var data []primitive.M
	for cur.Next(context.Background()) {
		var element primitive.M
		err := cur.Decode(&element)
		services.GetError(err)
		data = append(data, element)
	}
	// cur.All(context.TODO(), &data)

	if len(data) < 1 {
		fmt.Println("No data found")
	}

	fmt.Println()

	defer cur.Close(context.Background())
	defer r.Body.Close()
	json.NewEncoder(w).Encode(data)
}

func DeleteGpuData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")

	// Get the Collection
	collection := services.GetGpuCollection()

	// Find the data, run the Query
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{})

	services.GetError(err)

	fmt.Println("Deleted", deleteResult.DeletedCount, "documents in the Gpu collection")

	defer r.Body.Close()
}

func CreateGpuData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "POST")

	// Get the json received in the request body and store on the variable
	var data []model.Gpu
	err := json.NewDecoder(r.Body).Decode(&data)
	services.GetError(err)

	// Set the interface for the docs to be inserted
	docs := make([]interface{}, len(data))
	for i, u := range data {
		docs[i] = u
	}

	// Get the Collection
	collection := services.GetGpuCollection()

	// delete the previous data
	collection.DeleteMany(context.TODO(), bson.D{})
	fmt.Println("Previous data deleted")

	// Run the Query
	collection.InsertMany(context.TODO(), docs)

	fmt.Println("Inserted", len(docs), "documents in the Gpu collection")

	defer r.Body.Close()
}

// for _, filter := range filters {
// 	cur, err := collection.Find(context.TODO(), bson.D{{Key: filter.Parameter, Value: filter.Value}}, opts)
// 	services.GetError(err)
// 	defer cur.Close(context.TODO())
// 	var results []bson.M
// 	if err = cur.All(context.TODO(), &results); err != nil {
// 		services.GetError(err)
// 	}
// 	json.NewEncoder(w).Encode(results)
// }

// bson.M{{Key: "brand", Value: "nvidia"}},
// bson.M{{Key: "brand", Value: filters[0]}},

// filters contains all the filters you want to apply:
// filters := []bson.M{
// 	{"brand": parameters[1]},
// 	{"brand": parameters[0]},
// }

// for k, v := range parameters {
// 	// filters = append(filters, {k: v})
// 	fmt.Println(k, ":", v)
// }

// fmt.Println("filters: ", filters)

// filter := bson.M{
// 	"$and": []bson.M{ // you can try this in []interface
// 		{"$or": filters},
// 		{"$or": []bson.M{{"model": "6600"}, {"model": "6400"}}},
// 	},
// }
