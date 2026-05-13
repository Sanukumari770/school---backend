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

// ==========================
// CREATE PARENT
// ==========================

func CreateParent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("parents")

	var parent models.Parent

	err := json.NewDecoder(r.Body).Decode(&parent)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parent.ID = primitive.NewObjectID()

	now := time.Now()

	parent.CreatedAt = now
	parent.UpdatedAt = now

	_, err = collection.InsertOne(
		context.Background(),
		parent,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Parent Created Successfully",
		"data": parent,
	})
}

// add multiple parents 

func AddMultipleParents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("parents")

	var parents []models.Parent

	err := json.NewDecoder(r.Body).Decode(&parents)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(parents) == 0 {

		http.Error(w, "No parents found", http.StatusBadRequest)
		return
	}

	var docs []interface{}

	for i := range parents {

		parents[i].ID = primitive.NewObjectID()

		now := time.Now()

		parents[i].CreatedAt = now
		parents[i].UpdatedAt = now

		docs = append(docs, parents[i])
	}

	result, err := collection.InsertMany(
		context.Background(),
		docs,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Multiple Parents Added Successfully",
		"inserted_ids": result.InsertedIDs,
	})
}


// ==========================
// GET ALL PARENTS
// ==========================

func GetParents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("parents")

	filter := bson.M{
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	cursor, err := collection.Find(
		context.Background(),
		filter,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var parents []models.Parent

	err = cursor.All(context.Background(), &parents)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if parents == nil {

		parents = []models.Parent{}
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"count": len(parents),
		"data": parents,
	})
}

// ==========================
// GET FULL PARENT DATA
// ==========================

func GetParentFull(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("parents")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Parent ID", http.StatusBadRequest)
		return
	}

	pipeline := bson.A{

		bson.M{
			"$match": bson.M{
				"_id": objID,
			},
		},

		// STUDENTS
		bson.M{
			"$lookup": bson.M{
				"from": "students",
				"localField": "student_ids",
				"foreignField": "_id",
				"as": "students",
			},
		},

		// ATTENDANCE
		bson.M{
			"$lookup": bson.M{
				"from": "attendance",
				"localField": "student_ids",
				"foreignField": "student_id",
				"as": "attendance",
			},
		},

		// MARKS
		bson.M{
			"$lookup": bson.M{
				"from": "marks",
				"localField": "student_ids",
				"foreignField": "student_id",
				"as": "marks",
			},
		},

		// FEES
		bson.M{
			"$lookup": bson.M{
				"from": "fees",
				"localField": "student_ids",
				"foreignField": "student_id",
				"as": "fees",
			},
		},

		// TRANSPORT
		bson.M{
			"$lookup": bson.M{
				"from": "transport",
				"localField": "student_ids",
				"foreignField": "student_id",
				"as": "transport",
			},
		},

		// ASSIGNMENTS
		bson.M{
			"$lookup": bson.M{
				"from": "assignments",
				"localField": "student_ids",
				"foreignField": "student_id",
				"as": "assignments",
			},
		},
	}

	cursor, err := collection.Aggregate(
		context.Background(),
		pipeline,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []bson.M

	err = cursor.All(context.Background(), &result)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// ==========================
// UPDATE PARENT
// ==========================

func UpdateParent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("parents")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Parent ID", http.StatusBadRequest)
		return
	}

	var updateData bson.M

	err = json.NewDecoder(r.Body).Decode(&updateData)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateData["updated_at"] = time.Now()

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": updateData,
		},
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Parent Updated Successfully",
	})
}

// ==========================
// DELETE PARENT
// ==========================

func DeleteParent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	collection := config.DB.Collection("parents")

	id := mux.Vars(r)["id"]

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		http.Error(w, "Invalid Parent ID", http.StatusBadRequest)
		return
	}

	now := time.Now()

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": bson.M{
				"deleted_at": now,
			},
		},
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Parent Deleted Successfully",
	})
}