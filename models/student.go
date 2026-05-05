package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`// students name 
	ClassID  primitive.ObjectID `bson:"class_id"`// class id generate from create class 
	RollNo   string      `bson:"roll_no"`// students roll no 
	Email    string    `bson:"email"` // students email 
	Phone    string    `bson:"phone"`// students phone number 
	Class    string     `bson:"class"`     //  add class 
	Section  string    `bson:"section"`   //  add
CreatedAt time.Time  `bson:"createdAt"`
ParentID  primitive.ObjectID `bson:"parent_id,omitempty"` // parents id from create parents 
UpdatedAt time.Time  `bson:"updatedAt"` // students created time 
DeletedAt *time.Time `bson:"deletedAt,omitempty"`// studets deleted time 
}