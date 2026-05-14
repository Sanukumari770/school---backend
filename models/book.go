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

	TotalBooks int `json:"total_books" bson:"total_books"`

	AvailableBooks int `json:"available_books" bson:"available_books"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}