
//admission application form and application number 

package models
import (
	"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	ApplicationNo string `bson:"applicationNo" json:"applicationNo"`

	// parent relation
	ParentID primitive.ObjectID `bson:"parentId" json:"parentId"`

	// class relation
	ClassID primitive.ObjectID `bson:"classId" json:"classId"`

	// student info
	StudentName string `bson:"studentName" json:"studentName"`
	DOB string `bson:"dob" json:"dob"`

	Email string `bson:"email" json:"email"`
	Phone string `bson:"phone" json:"phone"`

	Section string `bson:"section" json:"section"`

	// documents
	Documents []Document `bson:"documents" json:"documents"`

	// entrance
	EntranceMarks int `bson:"entranceMarks" json:"entranceMarks"`
	EntranceResult string `bson:"entranceResult" json:"entranceResult"`

	// merit
	MeritRank int `bson:"meritRank" json:"meritRank"`

	// seat
	SeatAllocated bool `bson:"seatAllocated" json:"seatAllocated"`

	// hostel / transport
	Transport bool `bson:"transport" json:"transport"`
	Hostel bool `bson:"hostel" json:"hostel"`

	// payment
	FeePaid bool `bson:"feePaid" json:"feePaid"`

	// final status
	Status string `bson:"status" json:"status"`
	/*
		applied
		entrance_completed
		selected
		rejected
		admitted
	*/

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}