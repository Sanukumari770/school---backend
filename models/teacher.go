package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teacher struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Subject   string             `bson:"subject"`
	ClassIDs  []string           `bson:"classIds"`
}