package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Parent struct {

	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Name string `json:"name" bson:"name"`

	Email string `json:"email" bson:"email"`

	Password string `json:"password" bson:"password"`

	Phone string `json:"phone" bson:"phone"`

	StudentIDs []primitive.ObjectID `json:"student_ids" bson:"student_ids"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}