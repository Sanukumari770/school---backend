package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	ClassID  primitive.ObjectID `bson:"class_id"`
	RollNo   string             `bson:"roll_no"`
	Email    string             `bson:"email"`
	Phone    string             `bson:"phone"`

	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}