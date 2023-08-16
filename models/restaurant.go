package models

import(
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct{
    Building string `json:"building" bson:"building,required"`
    Coord []float64 `json:"coord" bson:"coord,required"`
    Street string `json:"street" bson:"street,required"`
    Zipcode string `json:"zipcode" bson:"zipcode,required"`
}

type Grades struct{
    Date time.Time  `json:"date" bson:"date,required"`
    Grade string `json:"grade" bson:"grade,required"`
    Score int `json:"score" bson:"score,required"`
}

type Restaurant struct{
    ID primitive.ObjectID `json:"id bson:"_id`
    Address1 Address `json:"address1" bson:"address1,required"`
    Borough string `json:"borough" bson:"borough,required"`
    Cuisine string `json:"cuisine" bson:"cuisine,required"`
    Grades1 []Grades `json:"grades" bson:"grades,required"`
    Name string `json:"name" bson:"name,required"`
    Restaurant_id string `json:"restaurant_id" bson:"estaurant_id,required"`

}

