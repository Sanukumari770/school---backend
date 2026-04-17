package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school/config"
	"school/models"
	"school/utils"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// ➕ Register
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "All fields required", 400)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config.DB.Collection("users").InsertOne(ctx, user)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Registered",
	})
}

// 🔐 Login
func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var found models.User
	err := config.DB.Collection("users").
		FindOne(ctx, bson.M{"email": user.Email}).
		Decode(&found)

	if err != nil {
		http.Error(w, "User not found", 401)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	token, _ := utils.GenerateToken(found.ID.Hex(), found.Role)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}