package api

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"shri-krishangan/config"
	"shri-krishangan/models"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductController struct {
	*mongo.Database
}

func NewProductController(db *mongo.Database) *ProductController {
	return &ProductController{
		Database: db,
	}
}

func (c *ProductController) ListProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		cur, err := c.Collection("products").Aggregate(config.DefaultCtx(), []bson.M{
			{"$skip": skip},
			{"$limit": limit},
		})

		if err != nil {
			log.Println(err)
			return
		}
		var products []models.Product
		err = cur.All(config.DefaultCtx(), &products)
		w.Header().Set("content-type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func (c *ProductController) AddProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		styles := readCsvFile("E:\\Quick Projects\\React_Project\\goserver\\dataset\\styles.csv")
		images := readCsvFile("E:\\Quick Projects\\React_Project\\goserver\\dataset\\images.csv")
		//var products []models.Product
		for i, style := range styles {
			var product models.Product
			product.Id = primitive.NewObjectIDFromTimestamp(time.Now())
			product.Name = style[9]
			product.Image = images[i][1]
			product.Price = 1000
			res, err := c.Database.Collection("products").InsertOne(config.DefaultCtx(), product)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("added to db :", res)
			//products = append(products, product)
		}
		res := "List added to DB"
		w.Write([]byte(res))
	}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
