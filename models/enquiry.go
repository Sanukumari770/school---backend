
package models

import (
"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Enquiry struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name   string `bson:"name"`
	Phone  string `bson:"phone"`
	Email  string `bson:"email"`
	Course string `bson:"course"`
	ClassInterested string `bson:"classInterested"`
Status string `bson:"status"` // new, contacted, converted
CreatedAt time.Time `bson:"createdAt"`
}