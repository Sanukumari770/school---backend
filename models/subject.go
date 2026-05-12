package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // created subjects id 
	Name      string             `bson:"name"`
	ClassID   primitive.ObjectID `bson:"class_id"`
	TeacherID primitive.ObjectID `bson:"teacher_id"`
// subjects created , deleted , updated timing mention through this stru 
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}