package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	ClassID   primitive.ObjectID `bson:"class_id"`
	TeacherID primitive.ObjectID `bson:"teacher_id"`

	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}