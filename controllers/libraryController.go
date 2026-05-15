package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ADD SINGLE BOOK

func AddBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	book.ID = primitive.NewObjectID()
	book.CreatedAt = time.Now()

	// agar available copies nahi bheja
	if book.AvailableCopies == 0 {
		book.AvailableCopies = book.TotalCopies
	}

	_, err = config.DB.Collection("books").InsertOne(
		context.TODO(),
		book,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Book Added Successfully",
	})
}

// ADD MULTIPLE BOOKS

func AddMultipleBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var books []models.Book

	err := json.NewDecoder(r.Body).Decode(&books)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var docs []interface{}

	for i := range books {

		books[i].ID = primitive.NewObjectID()
		books[i].CreatedAt = time.Now()

		if books[i].AvailableCopies == 0 {
			books[i].AvailableCopies = books[i].TotalCopies
		}

		docs = append(docs, books[i])
	}

	_, err = config.DB.Collection("books").InsertMany(
		context.TODO(),
		docs,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Multiple Books Added Successfully",
		"total_books_added": len(books),
	})
}

// GET ALL BOOKS

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cursor, err := config.DB.Collection("books").Find(
		context.TODO(),
		bson.M{},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer cursor.Close(context.TODO())

	var books []models.Book

	cursor.All(context.TODO(), &books)

	if books == nil {
		books = []models.Book{}
	}

	json.NewEncoder(w).Encode(books)
}

// ISSUE BOOK

func IssueBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var issue models.LibraryIssue

	err := json.NewDecoder(r.Body).Decode(&issue)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// CHECK BOOK EXISTS
	var book models.Book

	err = config.DB.Collection("books").FindOne(
		context.TODO(),
		bson.M{
			"_id": issue.BookID,
		},
	).Decode(&book)

	if err != nil {
		http.Error(w, "Book not found", 404)
		return
	}

	// CHECK AVAILABLE COPIES
	if book.AvailableCopies <= 0 {
		http.Error(w, "Book not available", 400)
		return
	}

	issue.ID = primitive.NewObjectID()
	issue.IssueDate = time.Now()
	issue.Status = "Issued"
	issue.CreatedAt = time.Now()

	_, err = config.DB.Collection("library_issue").InsertOne(
		context.TODO(),
		issue,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// DECREASE AVAILABLE COPIES
	_, err = config.DB.Collection("books").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": issue.BookID,
		},
		bson.M{
			"$inc": bson.M{
				"available_copies": -1,
			},
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Book Issued Successfully",
	})
}

// RETURN BOOK

func ReturnBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	var issue models.LibraryIssue

	err = config.DB.Collection("library_issue").FindOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
	).Decode(&issue)

	if err != nil {
		http.Error(w, "Issue not found", 404)
		return
	}

	// UPDATE ISSUE STATUS
	_, err = config.DB.Collection("library_issue").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": bson.M{
				"status":      "Returned",
				"return_date": time.Now(),
			},
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// INCREASE AVAILABLE COPIES
	_, err = config.DB.Collection("books").UpdateOne(
		context.TODO(),
		bson.M{
			"_id": issue.BookID,
		},
		bson.M{
			"$inc": bson.M{
				"available_copies": 1,
			},
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(bson.M{
		"success": true,
		"message": "Book Returned Successfully",
	})
}

// FULL LIBRARY DETAILS

func GetLibraryDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	pipeline := bson.A{

		// JOIN STUDENTS
		bson.M{
			"$lookup": bson.M{
				"from":         "students",
				"localField":   "student_id",
				"foreignField": "_id",
				"as":           "student",
			},
		},

		// JOIN BOOKS
		bson.M{
			"$lookup": bson.M{
				"from":         "books",
				"localField":   "book_id",
				"foreignField": "_id",
				"as":           "book",
			},
		},
	}

	cursor, err := config.DB.Collection("library_issue").Aggregate(
		context.TODO(),
		pipeline,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var result []bson.M

	cursor.All(context.TODO(), &result)

	if result == nil {
		result = []bson.M{}
	}

	json.NewEncoder(w).Encode(result)
}