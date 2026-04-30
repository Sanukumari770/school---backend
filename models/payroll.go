package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type Payroll struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TeacherID primitive.ObjectID `bson:"teacher_id"`

	Salary int    `bson:"salary"`
	Month  string `bson:"month"`

	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}