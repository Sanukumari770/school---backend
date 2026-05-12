package models
import (
"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Name string `json:"name" bson:"name"`

	ClassID *primitive.ObjectID `json:"class_id,omitempty" bson:"class_id,omitempty"`

	RollNo string `json:"roll_no" bson:"roll_no"`

	Email string `json:"email" bson:"email"`

	Phone string `json:"phone" bson:"phone"`

	Class string `json:"class" bson:"class"`

	Section string `json:"section" bson:"section"`

	CreatedAt time.Time `json:"created_at" bson:"createdAt"`

	UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"`

	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deletedAt,omitempty"`

	ParentID *primitive.ObjectID `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
}