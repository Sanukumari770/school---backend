package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `json:"name"`
	Date string             `json:"date"`
}