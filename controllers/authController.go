package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"school/config"
	"school/models"
	"school/utils"

	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	collection := config.DB.Collection("users")
	collection.InsertOne(context.TODO(), user)

	json.NewEncoder(w).Encode("Registered")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	json.NewDecoder(r.Body).Decode(&input)

	collection := config.DB.Collection("users")
	collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Invalid", 401)
		return
	}

	token, _ := utils.GenerateJWT(user.ID.Hex(), user.Role)

	json.NewEncoder(w).Encode(token)
}