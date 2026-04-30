package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Timetable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"class_id"`
	SubjectID primitive.ObjectID `bson:"subject_id"`
	TeacherID primitive.ObjectID `bson:"teacher_id"`

	TimeSlot string `bson:"time_slot"`

	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}