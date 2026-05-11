package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bus struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	BusNo string `json:"bus_no" bson:"bus_no"`

	Route string `json:"route" bson:"route"`

	DriverName string `json:"driver_name" bson:"driver_name"`

	DriverPhone string `json:"driver_phone" bson:"driver_phone"`

	TotalSeats int `json:"total_seats" bson:"total_seats"`

	OccupiedSeats int `json:"occupied_seats" bson:"occupied_seats"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}