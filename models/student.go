package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	ClassID string             `bson:"classId"`
	RollNo  string             `bson:"rollNo"`
	Email   string             `bson:"email"`
	Phone   string             `bson:"phone"`
	CreatedAt time.Time        `bson:"createdAt"`
}