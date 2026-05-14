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

	Qualification string `json:"qualification" bson:"qualification"`

	Experience string `json:"experience" bson:"experience"`

	//  this struc for time mention of created and deleted , updated 

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`

    UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`

    DeletedAt *time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}