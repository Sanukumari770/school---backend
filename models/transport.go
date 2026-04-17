package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transport struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID string             `json:"studentId"`
	BusID     string             `json:"busId"`
}