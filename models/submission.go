package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	AssignmentID primitive.ObjectID `bson:"assignment_id"`
	StudentID    primitive.ObjectID `bson:"student_id"`

	FileURL string `bson:"file_url"`
	Marks   int    `bson:"marks"`

	SubmittedAt time.Time `bson:"submittedAt"`
}