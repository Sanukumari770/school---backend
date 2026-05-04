package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payroll struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TeacherID primitive.ObjectID `bson:"teacher_id"`

	Salary    float64 `bson:"salary"`
	Bonus     float64 `bson:"bonus"`
	Deduction float64 `bson:"deduction"`

	Month string `bson:"month"`

	CreatedAt time.Time `bson:"createdAt"`
}