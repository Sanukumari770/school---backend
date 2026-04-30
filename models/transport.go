package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transport struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID primitive.ObjectID `bson:"student_id"`

	BusNo string `bson:"bus_no"`
	Route string `bson:"route"`
}