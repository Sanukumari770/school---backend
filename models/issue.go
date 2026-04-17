package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Issue struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BookID    string             `bson:"bookId"`
	StudentID string             `bson:"studentId"`
	ReturnDate string            `bson:"returnDate"`
}