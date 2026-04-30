package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receipt struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	StudentID primitive.ObjectID `bson:"student_id"`
	PaymentID primitive.ObjectID `bson:"payment_id"`

	ReceiptNo string `bson:"receipt_no"`

	Amount float64 `bson:"amount"`

	CreatedAt time.Time `bson:"createdAt"`
}