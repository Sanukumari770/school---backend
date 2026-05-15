package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exam struct {

	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	ExamName string `json:"exam_name" bson:"exam_name"`

	ClassName string `json:"class_name" bson:"class_name"`

	Subject string `json:"subject" bson:"subject"`

	ExamDate string `json:"exam_date" bson:"exam_date"`

	MaxMarks int `json:"max_marks" bson:"max_marks"`

	PassMark int `json:"pass_mark" bson:"pass_mark"`

	Status string `json:"status" bson:"status"`
	
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}