package models
import (
"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // createted  students id from api testing 

	Name string `json:"name" bson:"name"` // name of students 

	ClassID *primitive.ObjectID `json:"class_id,omitempty" bson:"class_id,omitempty"`// fetch class id from class api test

	RollNo string `json:"roll_no" bson:"roll_no"` // roll number of students 

	Email string `json:"email" bson:"email"` // email of students 

	Phone string `json:"phone" bson:"phone"` // phone number of students 

	Class string `json:"class" bson:"class"` // class of students 

	Section string `json:"section" bson:"section"` // section with class 

	FatherName string `json:"father_name" bson:"father_name"`

	FatherEmail string `json:"father_email" bson:"father_email"`

	CreatedAt time.Time `json:"created_at" bson:"createdAt"`  // students adding time 

	UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"` // udated deatils time 

	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deletedAt,omitempty"`// delet 

	ParentID *primitive.ObjectID `json:"parent_id,omitempty" bson:"parent_id,omitempty"`  // fetch parents details fromm id 
}