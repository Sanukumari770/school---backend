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


// =======================
// ADD SINGLE STUDENT
// =======================

func AddStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var input struct {
		Name    string `json:"name"`
		ClassID string `json:"class_id"`
		RollNo  string `json:"roll_no"`
		Class   string `json:"class"`
		Section string `json:"section"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var classObjID *primitive.ObjectID

	if input.ClassID != "" {

		objID, err := primitive.ObjectIDFromHex(input.ClassID)
		if err == nil {
			classObjID = &objID
		}
	}

	student := models.Student{
		ID:        primitive.NewObjectID(),
		Name:      input.Name,
		ClassID:   classObjID,
		RollNo:    input.RollNo,
		Class:     input.Class,
		Section:   input.Section,
		Email:     input.Email,
		Phone:     input.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = config.DB.Collection("students").InsertOne(
		context.TODO(),
		student,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Student Added Successfully",
	})
}


// =======================
// ADD MULTIPLE STUDENTS
// =======================

func AddMultipleStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var input []struct {
		Name    string `json:"name"`
		ClassID string `json:"class_id"`
		RollNo  string `json:"roll_no"`
		Class   string `json:"class"`
		Section string `json:"section"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var docs []interface{}

	for _, s := range input {

		var classObjID *primitive.ObjectID

		if s.ClassID != "" {

			objID, err := primitive.ObjectIDFromHex(s.ClassID)
			if err == nil {
				classObjID = &objID
			}
		}

		student := models.Student{
			ID:        primitive.NewObjectID(),
			Name:      s.Name,
			ClassID:   classObjID,
			RollNo:    s.RollNo,
			Class:     s.Class,
			Section:   s.Section,
			Email:     s.Email,
			Phone:     s.Phone,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		docs = append(docs, student)
	}

	if len(docs) == 0 {
		http.Error(w, "No valid students found", http.StatusBadRequest)
		return
	}

	result, err := config.DB.Collection("students").InsertMany(
		context.TODO(),
		docs,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success":      true,
		"inserted_ids": result.InsertedIDs,
	})
}


// =======================
// GET ALL STUDENTS
// =======================

func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	filter := bson.M{
		"deletedAt": bson.M{
			"$exists": false,
		},
	}

	cursor, err := config.DB.Collection("students").Find(
		context.TODO(),
		filter,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	var students []models.Student

	err = cursor.All(context.TODO(), &students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if students == nil {
		students = []models.Student{}
	}

	json.NewEncoder(w).Encode(students)
}


// =======================
// GET SINGLE STUDENT
// =======================

func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	var student models.Student

	err = config.DB.Collection("students").FindOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
	).Decode(&student)

	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}


// =======================
// GET FULL STUDENT DETAILS
// =======================

func GetStudentFull(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	pipeline := bson.A{

		bson.M{
			"$match": bson.M{
				"_id": objID,
			},
		},

		// CLASS
		bson.M{
			"$lookup": bson.M{
				"from":         "classes",
				"localField":   "class_id",
				"foreignField": "_id",
				"as":           "class_data",
			},
		},

		// SUBJECTS
		bson.M{
			"$lookup": bson.M{
				"from":         "subjects",
				"localField":   "class_id",
				"foreignField": "class_id",
				"as":           "subjects",
			},
		},

		// ATTENDANCE
		bson.M{
			"$lookup": bson.M{
				"from":         "attendance",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "attendance",
			},
		},

		// MARKS
		bson.M{
			"$lookup": bson.M{
				"from":         "marks",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "marks",
			},
		},

		// FEES
		bson.M{
			"$lookup": bson.M{
				"from":         "fees",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "fees",
			},
		},

		// PARENT
		bson.M{
			"$lookup": bson.M{
				"from":         "parents",
				"localField":   "parent_id",
				"foreignField": "_id",
				"as":           "parent",
			},
		},

		// ASSIGNMENTS
		bson.M{
			"$lookup": bson.M{
				"from":         "assignments",
				"localField":   "class_id",
				"foreignField": "class_id",
				"as":           "assignments",
			},
		},

		// SUBMISSIONS
		bson.M{
			"$lookup": bson.M{
				"from":         "submissions",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "submissions",
			},
		},

		// TRANSPORT
		bson.M{
			"$lookup": bson.M{
				"from":         "transport",
				"localField":   "_id",
				"foreignField": "student_id",
				"as":           "transport",
			},
		},
	}

	cursor, err := config.DB.Collection("students").Aggregate(
		context.TODO(),
		pipeline,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []bson.M

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}


// =======================
// UPDATE STUDENT
// =======================

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	var update bson.M

	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	update["updatedAt"] = time.Now()

	_, err = config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": update,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Student Updated Successfully",
	})
}


// =======================
// DELETE STUDENT
// =======================

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	now := time.Now()

	_, err = config.DB.Collection("students").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": bson.M{
				"deletedAt": now,
			},
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Student Deleted Successfully",
	})
}