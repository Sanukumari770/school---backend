package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	ClassID   primitive.ObjectID `bson:"class_id" json:"class_id"`
	TeacherID primitive.ObjectID `bson:"teacher_id" json:"teacher_id"`

	Date string `bson:"date" json:"date"`

	Marks  int    `bson:"marks" json:"marks"`   // add
	Result string `bson:"result" json:"result"` // add

	CreatedAt time.Time  `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt" json:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}