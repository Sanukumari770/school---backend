package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Marks struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID primitive.ObjectID `bson:"student_id"`
	ExamID    primitive.ObjectID `bson:"exam_id"`
	SubjectID primitive.ObjectID `bson:"subject_id"`

	Marks int    `bson:"marks"`
	Grade string `bson:"grade"`

	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}