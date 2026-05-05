
package controllers
import (
"context"
"encoding/json"
"net/http"
"time"
"school/config"
"school/models"
"github.com/gorilla/mux"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
)

//  Create Student
func AddMultipleStudents(w http.ResponseWriter, r *http.Request) {

	var input []struct {
		Name    string `json:"name"`
		ClassID string `json:"class_id"`
		RollNo  string `json:"roll_no"`
		Class   string `json:"class"`
		Section string `json:"section"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	var docs []interface{}

	for _, s := range input {

		classID, _ := primitive.ObjectIDFromHex(s.ClassID)

		data := models.Student{
			Name:      s.Name,
			ClassID:   classID,
			RollNo:    s.RollNo,
			Class:     s.Class,
			Section:   s.Section,
			Email:     s.Email,
			Phone:     s.Phone,
			CreatedAt: time.Now(),
		}

		docs = append(docs, data)
	}

	res, _ := config.DB.Collection("students").InsertMany(context.TODO(), docs)

	json.NewEncoder(w).Encode(res)
}

// get students fetch data from create students 
func GetStudents(w http.ResponseWriter, r *http.Request) {

	cursor, err := config.DB.Collection("students").Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var students []models.Student

	if err := cursor.All(context.TODO(), &students); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if students == nil {
		students = []models.Student{} // avoid null
	}

	json.NewEncoder(w).Encode(students)
}
// Get Full Student (JOIN)
func GetStudentFull(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	pipeline := bson.A{

		// student select 
		bson.M{"$match": bson.M{"_id": objID}},

		// class
		bson.M{"$lookup": bson.M{
			"from": "classes",
			"localField": "class_id",
			"foreignField": "_id",
			"as": "class",
		}},

		// subjects
		bson.M{"$lookup": bson.M{
			"from": "subjects",
			"localField": "class_id",
			"foreignField": "class_id",
			"as": "subjects",
		}},

		// attendance 
		bson.M{"$lookup": bson.M{
			"from": "attendance",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "attendance",
		}},

		// marks 
		bson.M{"$lookup": bson.M{
			"from": "marks",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "marks",
		}},

		// exam 

		bson.M{"$lookup": bson.M{
			"from": "timetable",
			"localField": "class_id",
			"foreignField": "class_id",
			"as": "timetable",
		}},


		// fees
		bson.M{"$lookup": bson.M{
			"from": "fees",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "fees",
		}},

		// books 
		bson.M{"$lookup": bson.M{
			"from": "books",
			"localField": "class_id",
			"foreignField": "class",
			"as": "books",
		}},

		//  Parent
		bson.M{"$lookup": bson.M{
			"from": "parents",
			"localField": "parent_id",
			"foreignField": "_id",
			"as": "parent",
		}},

		//  Assignments
		bson.M{"$lookup": bson.M{
			"from": "assignments",
			"localField": "class_id",
			"foreignField": "class_id",
			"as": "assignments",
		}},

		//  Submissions
		bson.M{"$lookup": bson.M{
			"from": "submissions",
			"localField": "_id",
			"foreignField": "student_id",
			"as": "submissions",
		}},
	}

	cursor, _ := config.DB.Collection("students").Aggregate(context.TODO(), pipeline)

	var result []bson.M
	cursor.All(context.TODO(), &result)

	json.NewEncoder(w).Encode(result)
}

//  Update Student
func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	// take id from URL 
	id := mux.Vars(r)["id"]

	// objectid convert 
	objID, _ := primitive.ObjectIDFromHex(id)

	// take data from body 
	var update bson.M
	json.NewDecoder(r.Body).Decode(&update)

	// add updateAt 
	update["updatedAt"] = time.Now()

	// mongodb update 
	config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": update},
	)


	// response 
	json.NewEncoder(w).Encode("Updated")
}

//  Delete Student (Soft Delete)
func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	now := time.Now()

	config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"deletedAt": now}},
	)

	json.NewEncoder(w).Encode("Deleted")
}