package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {

	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Title string `json:"title" bson:"title"`

	Subject string `json:"subject" bson:"subject"`

	ClassName string `json:"className" bson:"className"`

	TeacherName string `json:"teacherName" bson:"teacherName"`

	DueDate string `json:"dueDate" bson:"dueDate"`

	TotalStudents int `json:"totalStudents" bson:"totalStudents"`

	SubmittedCount int `json:"submittedCount" bson:"submittedCount"`

	Status string `json:"status" bson:"status"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Submission struct {

	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	AssignmentID primitive.ObjectID `json:"assignment_id" bson:"assignment_id"`

	StudentName string `json:"student_name" bson:"student_name"`

	FileURL string `json:"file_url" bson:"file_url"`

	Marks int `json:"marks" bson:"marks"`

	SubmittedAt time.Time `json:"submittedAt" bson:"submittedAt"`
}