package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hostel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID primitive.ObjectID `bson:"student_id"`

	RoomNo string `bson:"room_no"`
	Block  string `bson:"block"`
}