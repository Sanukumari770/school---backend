package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admission struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	ApplicationNo string `bson:"applicationNo" json:"applicationNo"`

	// 🔗 relation
	ParentID primitive.ObjectID `bson:"parentId" json:"parentId"`
	ClassID  primitive.ObjectID `bson:"classId" json:"classId"`

	// student info
	StudentName string `bson:"studentName" json:"studentName"`
	Email       string `bson:"email" json:"email"`
	Phone       string `bson:"phone" json:"phone"`
	Section     string `bson:"section" json:"section"`

	// documents
	Documents []Document `bson:"documents" json:"documents"`

	// entrance
	EntranceScore int    `bson:"entranceScore,omitempty" json:"entranceScore"`
	MeritRank     int    `bson:"meritRank,omitempty" json:"meritRank"`
	Result        string `bson:"result" json:"result"` // pass/fail

	// facilities
	Transport bool `bson:"transport" json:"transport"`
	Hostel    bool `bson:"hostel" json:"hostel"`

	// payment
	FeePaid bool `bson:"feePaid" json:"feePaid"`

	// status
	Status string `bson:"status" json:"status"` // pending/approved/rejected

	CreatedAt time.Time  `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt" json:"updatedAt"`
}