package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	SubjectID primitive.ObjectID `bson:"subject_id"`
	TeacherID primitive.ObjectID `bson:"teacher_id"`
	ClassID   primitive.ObjectID `bson:"class_id"`

	DueDate   time.Time `bson:"due_date"`

	CreatedAt time.Time `bson:"createdAt"`
}