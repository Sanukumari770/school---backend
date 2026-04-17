package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Timetable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   string             `json:"classId"`
	SubjectID string             `json:"subjectId"`
	TeacherID string             `json:"teacherId"`
	TimeSlot  string             `json:"timeSlot"`
}