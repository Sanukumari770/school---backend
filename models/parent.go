package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Parent struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone string             `bson:"phone"`

	StudentIDs []primitive.ObjectID `bson:"student_ids"`

	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}