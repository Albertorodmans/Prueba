package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transaccion struct {
	TransaccionId bson.ObjectId `json:"transaccion_id" bson:"transaccion_id"`
	MerchantId    bson.ObjectId `json:"merchant_id" bson:"merchant_id"`
	Amount        float64       `json:"amount" bson:"amount"`
	Commission    int           `json:"commission" bson:"commission"`
	Fee           string        `json:"fee" bson:"fee"`
	CreatedAt     time.Time     `json:"created_at" bson:"created_at"`
}

type Transaccions []Transaccion
