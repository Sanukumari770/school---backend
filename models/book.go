package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Title string `json:"title" bson:"title"`

	Author string `json:"author" bson:"author"`

	ISBN string `json:"isbn" bson:"isbn"`

	Category string `json:"category" bson:"category"`

	TotalCopies int `json:"total_copies" bson:"total_copies"`

	AvailableCopies int `json:"available_copies" bson:"available_copies"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}