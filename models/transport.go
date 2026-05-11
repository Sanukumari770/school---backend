package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport struct {

ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
StudentID primitive.ObjectID `json:"student_id" bson:"student_id"`// students who prefer bus 
BusID primitive.ObjectID `json:"bus_id" bson:"bus_id"`// bus id no 
BusNo string `json:"bus_no" bson:"bus_no"`// bus no 
Route string `json:"route" bson:"route"`// oute from where students go 
CreatedAt time.Time `json:"createdAt" bson:"createdAt"` 
}