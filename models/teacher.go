package models

import (
	"time"
// this stru for bson and time define 
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // created id 

	Name string `json:"name" bson:"name"` // name of teachers 

	Email string `json:"email" bson:"email"` // teachers email 

	Subject string `json:"subject" bson:"subject"` // assign subjects of teacher 

	Class string `json:"class" bson:"class"` // assign class fetch from class id 

	//  this struc for time mention of created and deleted , updated 
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}