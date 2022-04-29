package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Comercio struct {
	MerchantId   bson.ObjectId `json:"merchant_id" bson:"merchant_id"`
	MerchantName string        `json:"merchant_name" bson:"merchant_name"`
	Commission   int           `json:"commission" bson:"commission"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
}

type Comercios []*Comercio
