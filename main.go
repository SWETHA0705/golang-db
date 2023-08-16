package main

import (
	"fmt"
	//"golang-db/models"
	"golang-db/services"

	//"golang-db/config"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
)

func main() {
	//client,_:= config.ConnectDataBase()
	// collection :=config.GetCollection(client,"sample_training","zips")
	fmt.Println("mongodb successfully conected")
	//products := []interface{}{
	//	models.Product{ID: primitive.NewObjectID(), Name: "oneplus", Price: 100000, Description: "budget phone"},
	//	models.Product{ID: primitive.NewObjectID(), Name: "vivo", Price: 10000, Description: "china phone"}}
	//services.InsertProductList(products)
	res, _ := services.Findres()
	fmt.Println(res)
	for _, r := range res {
		fmt.Println(r)
	}

}
