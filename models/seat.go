package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seat struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	ApplicationID primitive.ObjectID `bson:"applicationId" json:"applicationId"`

	Class string `bson:"class" json:"class"`
	Section string `bson:"section" json:"section"`

	Status string `bson:"status" json:"status"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}