package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Name string `json:"name" bson:"name"`

	Email string `json:"email" bson:"email"`

	Subject string `json:"subject" bson:"subject"`

	Class string `json:"class" bson:"class"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`

	DeletedAt *time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}