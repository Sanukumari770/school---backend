// admission test 

package models

import (
"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)
type Entrance struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	AdmissionID primitive.ObjectID `bson:"admissionId"`

	Marks int `bson:"marks"`// marks obtain in exam 
	Result string `bson:"result"` // pass/fail
CreatedAt time.Time `bson:"createdAt"`
}