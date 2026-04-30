package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Type       string             `bson:"type" json:"type"`
	FileURL    string             `bson:"file_url" json:"file_url"`
	StudentID  primitive.ObjectID `bson:"student_id" json:"student_id"`
	UploadedAt time.Time          `bson:"uploaded_at" json:"uploaded_at"`
}