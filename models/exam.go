package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exam struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	SubjectID string             `bson:"subjectId"`
	TotalMarks int               `bson:"totalMarks"`
	ClassID primitive.ObjectID `bson:"class_id"`
	Date    string             `bson:"date"`
}