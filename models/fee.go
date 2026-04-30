package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fee struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	StudentID primitive.ObjectID `bson:"student_id"`
	ClassID   primitive.ObjectID `bson:"class_id"`

	TotalAmount float64 `bson:"total_amount"`
	PaidAmount  float64 `bson:"paid_amount"`
	DueAmount   float64 `bson:"due_amount"`

	Status string `bson:"status"` // paid / partial / unpaid

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}