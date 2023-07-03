package main

import (
	"log"
	"net/http"
	"shri-krishangan/api"
	"shri-krishangan/config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	db := config.DefaultDB()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	products := api.NewProductController(db)
	r.HandleFunc("/products", products.ListProducts()).Methods("GET")
	r.HandleFunc("/getlist", products.AddProducts()).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
