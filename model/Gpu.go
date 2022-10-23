package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gpu struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name"`
	Price      string             `json:"price"`
	PriceInt   int                `json:"priceInt"`
	Model      string             `json:"model"`
	Serie      string             `json:"serie"`
	Manufactor string             `json:"manufactor"`
	Brand      string             `json:"brand"`
	Store      string             `json:"store"`
	Link       string             `json:"link"`
}
