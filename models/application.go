
//admission application form and application number 

package models
import (
	"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	ApplicationNo string `bson:"applicationNo"`

	Name  string `bson:"name"`
	DOB   string `bson:"dob"`
	Class string `bson:"class"`
ParentName string `bson:"parentName"`
Phone      string `bson:"phone"`
ClassApplied string `bson:"classApplied"`
Status string `bson:"status"` // pending, approved
CreatedAt time.Time `bson:"createdAt"`
}