package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Enquiry struct {

	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Name string `bson:"name" json:"name"`

	Phone string `bson:"phone" json:"phone"`

	Email string `bson:"email" json:"email"`

	ClassInterested string `bson:"classInterested" json:"classInterested"`

	Message string `bson:"message" json:"message"`

	Status string `bson:"status" json:"status"`
	/*
		new
		contacted
		converted
		closed
	*/

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`

	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}