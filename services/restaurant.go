package services

import (
	"context"
	"fmt"
	"golang-db/config"
	"golang-db/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Rescontext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "sample_restaurants", "restaurants")
}

func InsertRestaurant(restaurant models.Restaurant) (*mongo.InsertOneResult, error) {
	cxt, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := Rescontext().InsertOne(cxt, restaurant)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}

func InsertRestaurantList(restaurant []interface{}) (*mongo.InsertManyResult, error) {
	cxt, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := Rescontext().InsertMany(cxt, restaurant)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}

func Findres() ([]*models.Restaurant, error) {
	cxt, _ := context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{}
	result, err := Rescontext().Find(cxt, filter, options.Find().SetLimit(10))
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		// storing our results in array
		var products []*models.Restaurant
		// using next it will check and iterate until it finds empty
		for result.Next(cxt) {
			// storing each output in product
			product := &models.Restaurant{}
			err := result.Decode(product)
			if err != nil {
				return nil, err

			}
			// using append we are adding product to our array products
			products = append(products, product)

		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {

			return []*models.Restaurant{}, nil
		}
		// we are returning our array products

		return products, nil

	}
}
