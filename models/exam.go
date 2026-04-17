package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exam struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `json:"name"`
	SubjectID string             `json:"subjectId"`
	TotalMarks int               `json:"totalMarks"`
}