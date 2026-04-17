package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Fees struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID string             `bson:"studentId"`
	Amount    int                `bson:"amount"`
	Status    string             `bson:"status"`
}