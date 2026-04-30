// merit list for admission 

package models

import (
"time"	
"go.mongodb.org/mongo-driver/bson/primitive"
)
type Merit struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	AdmissionID primitive.ObjectID `bson:"admissionId"`
	Rank int `bson:"rank"`
	CreatedAt time.Time `bson:"createdAt"`
}