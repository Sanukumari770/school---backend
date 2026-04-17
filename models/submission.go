package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Submission struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	AssignmentID string             `json:"assignmentId"`
	StudentID    string             `json:"studentId"`
	FileURL      string             `json:"fileUrl"`
	Marks        int                `json:"marks"`
}