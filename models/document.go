package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID string             `json:"studentId"`
	Type      string             `json:"type"`
	FileURL   string             `json:"fileUrl"`
}