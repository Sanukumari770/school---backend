package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type Bus struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Route string             `bson:"route"`
}