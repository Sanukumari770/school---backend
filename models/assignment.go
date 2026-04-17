package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Assignment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `json:"title"`
	SubjectID string             `json:"subjectId"`
	DueDate   string             `json:"dueDate"`
}