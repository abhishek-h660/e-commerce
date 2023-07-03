package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Price float32            `json:"price" bson:"price"`
	Image string             `json:"image" bson:"image"`
}
