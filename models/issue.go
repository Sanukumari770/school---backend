package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibraryIssue struct {

	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	StudentID primitive.ObjectID `json:"student_id" bson:"student_id"`

	BookID primitive.ObjectID `json:"book_id" bson:"book_id"`

	IssueDate time.Time `json:"issue_date" bson:"issue_date"`

	ReturnDate time.Time `json:"return_date" bson:"return_date"`

	Status string `json:"status" bson:"status"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}