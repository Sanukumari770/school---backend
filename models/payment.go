package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	StudentID primitive.ObjectID `bson:"student_id"`
	FeeID     primitive.ObjectID `bson:"fee_id"`

	Amount float64 `bson:"amount"`

	Method string `bson:"method"` // cash / online
	Status string `bson:"status"` // success / failed

	TransactionID string `bson:"transaction_id"`

	CreatedAt time.Time `bson:"createdAt"`
}