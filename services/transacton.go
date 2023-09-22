package services

import (
	"context"
	//"encoding/json"
	"fmt"

	"golang-db/config"
	//"golang-db/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

func TrnsContext()*mongo.Collection{
	//connecting the database 
client, _ := config.ConnectDataBase()
// geting collection from sample analytics
return config.GetCollection(client,"sample_analytics","transactions")

}



//  func FetchAggregatedTransactions(){
// 	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
// 	matchStage := bson.D{{"$match",bson.D{{"account_id",278866}}}}
// 	groupStage:= bson.D{
// 		{
// 			Key: "$group",Value: bson.D{
// 				{"_id", "$account_id"},
// 				{"total_count",bson.D{{"$sum","$transaction_count"}}},

// 			}}}
// 	result,err:= TrnsContext().Aggregate(ctx,mongo.Pipeline{matchStage,groupStage})
// 	if err!= nil{
// 		fmt.Println(err.Error)
		
// 	}else{
// 		var showsWithInfo [] bson.M
// 		if err = result.All(ctx,&showsWithInfo);err!=nil{
// 			panic(err)
// 		}
// 		formatted_data,err:= json.MarshalIndent(showsWithInfo,""," ")
// 		if err!= nil{
// 			fmt.Println(err.Error())
// 		}else{
// 			fmt.Println(string(formatted_data))
// 		}
// 	}
//  }
func UpdateTransaction(initialValue int,newValue int)(*mongo.UpdateResult,error){
	ctx,_ := context.WithTimeout(context.Background(),100*time.Second)
	filter:= bson.D{{"account_id",initialValue}}
	update:= bson.D{{"$Set",bson.D{{"account_id",newValue}}}}
	result,err:= TrnsContext().UpdateOne(ctx,filter,update)
	if err!= nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return result,nil
}