package models

import "time"

type Notice struct {
	ID          int       `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`// title for notice 
	Content     string    `json:"content" bson:"content"`// type of content 
	Priority    string    `json:"priority" bson:"priority"` // low, medium, high
	Category    string    `json:"category" bson:"category"`
	Audience    string    `json:"audience" bson:"audience"` // all, students, teachers, parents
	PublishDate string    `json:"publishDate" bson:"publishDate"`
	CreatedBy   string    `json:"createdBy" bson:"createdBy"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}