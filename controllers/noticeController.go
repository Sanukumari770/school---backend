package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"school/config"
	"school/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getNoticeCollection() *mongo.Collection {
	return config.GetCollection("notices")
}

// CREATE NOTICE
func CreateNotice(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var notice models.Notice

	err := json.NewDecoder(r.Body).Decode(&notice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notice.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// FIXED HERE
	result, err := getNoticeCollection().InsertOne(ctx, notice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Notice created successfully",
		"data":    result,
	})
}

// multiple notice 

func AddMultipleNotices(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var notices []models.Notice

	err := json.NewDecoder(r.Body).Decode(&notices)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var noticeList []interface{}

	for i := range notices {

		notices[i].CreatedAt = time.Now()

		noticeList = append(noticeList, notices[i])
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := getNoticeCollection().InsertMany(ctx, noticeList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Multiple notices added successfully",
		"data":    result.InsertedIDs,
	})
}

// GET ALL NOTICES
func GetNotices(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// FIXED HERE
	cursor, err := getNoticeCollection().Find(ctx, bson.M{})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var notices []models.Notice

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var notice models.Notice

		cursor.Decode(&notice)

		notices = append(notices, notice)
	}

	json.NewEncoder(w).Encode(notices)
}

// GET SINGLE NOTICE
func GetNotice(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

	// FIXED HERE
	noticeID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notice models.Notice

	// FIXED HERE
	err = getNoticeCollection().
		FindOne(ctx, bson.M{"id": noticeID}).
		Decode(&notice)

	if err != nil {
		http.Error(w, "Notice not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(notice)
}

// DELETE NOTICE
func DeleteNotice(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

	// FIXED HERE
	noticeID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// FIXED HERE
	_, err = getNoticeCollection().
		DeleteOne(ctx, bson.M{"id": noticeID})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Notice deleted successfully",
	})
}