package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attendance struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID string             `json:"studentId"`
	Date      string             `json:"date"`
	Status    string             `json:"status"`
}