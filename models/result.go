package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Result struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID string             `bson:"studentId"`
	ExamID    string             `bson:"examId"`
	Marks     int                `bson:"marks"`
	Grade     string             `bson:"grade"`
}