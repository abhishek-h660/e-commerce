package api

import (
	"encoding/json"
	"log"
	"net/http"
	"shri-krishangan/config"
	"shri-krishangan/models"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartController struct {
	*mongo.Database
}

func NewCartController(db *mongo.Database) *CartController {
	return &CartController{
		Database: db,
	}
}

func (c *CartController) ListCartItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		cur, err := models.CollectionCart(c.Database).Aggregate(config.DefaultCtx(), []bson.M{
			{"$skip": skip},
			{"$limit": limit},
			{"$sort": bson.M{"timestamp": -1}},
		})
		w.Header().Set("content-type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var cartItems []models.CartItem
		cur.All(config.DefaultCtx(), &cartItems)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cartItems)
	}
}

func (c *CartController) AddToCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cartItem models.CartItem
		w.Header().Set("content-type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&cartItem)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte(err.Error()))
			return
		}
		res, err := cartItem.Save(c.Database)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func (c *CartController) RemoveFromCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cartItem models.CartItem
		id := mux.Vars(r)["id"]
		_id, err := primitive.ObjectIDFromHex(id)
		cartItem.Id = _id

		cartItem.Delete(c.Database)
		w.Header().Set("content-type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Item Removed Successfully"))

	}
}
