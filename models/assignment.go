package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Assignment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	SubjectID string             `bson:"subjectId"`
	DueDate   string             `bson:"dueDate"`
}