package models

import (
	"log"
	"shri-krishangan/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartItem struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Product   Product            `json:"product" bson:"product"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

func CollectionCart(db *mongo.Database) *mongo.Collection {
	return db.Collection("cart")
}

func (c *CartItem) Save(db *mongo.Database) *mongo.InsertOneResult {
	res, err := CollectionCart(db).InsertOne(config.DefaultCtx(), c)
	if err != nil {
		log.Println(err)
	}
	return res
}

func (c *CartItem) Delete(db *mongo.Database) *mongo.DeleteResult {
	res, err := CollectionCart(db).DeleteOne(config.DefaultCtx(), bson.M{"_id": c.Id})
	if err != nil {
		log.Println(err)

	}
	return res
}
