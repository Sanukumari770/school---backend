package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hostel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID string             `json:"studentId"`
	RoomID    string             `json:"roomId"`
}