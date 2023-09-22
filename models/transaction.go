package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	AccountId int64 `json:"acoount_id" bson:"account_id"`
	TransactionCount int64 `json :"transactionCount" bson:"transaction_count,required"`
	Bucket_start_date time.Time `json:"startDate" bson:"bucket_start_date,required"`
	Bucket_end_date time.Time `json:"endDate" bson:"bucket_end_date"`
	Transactions []InnerTransaction `json:"innerTransaction" bson:"transactions"`

}
type InnerTransaction struct{
	Date time.Time `json:"date" bson:"date"`
	Amount int `json:"amount" bson:"amount"`
	Transactioncode string `json:"transaction_code" bson:"transaction_code,required"`
	Symbol string `json:"symbol" bson:"symbol"`
	Price string `json:"price" bson:"price"`
	
	Total string`json:"total" bson:"total"`
}

