package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport struct {

	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	StudentID primitive.ObjectID `json:"student_id" bson:"student_id"` // connected with students 

	ParentID primitive.ObjectID `json:"parent_id" bson:"parent_id"`  // connected with parents 

	BusID primitive.ObjectID `json:"bus_id" bson:"bus_id"`

	BusNo string `json:"bus_no" bson:"bus_no"`

	Route string `json:"route" bson:"route"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}