package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	ClassID primitive.ObjectID `bson:"class_id"` 
}