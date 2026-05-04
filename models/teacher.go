package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Teacher struct {
ID       primitive.ObjectID   `bson:"_id,omitempty"`
Name     string               `bson:"name"`
Email    string               `bson:"email"`
Subjects []primitive.ObjectID `bson:"subjects"` // which subject teaches
Classes  []primitive.ObjectID `bson:"classes"`
CreatedAt time.Time  `bson:"createdAt"`
UpdatedAt time.Time  `bson:"updatedAt"`
DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}