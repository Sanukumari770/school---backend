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


// ================= REGISTER =================
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	// 1️⃣ Decode request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// 2️⃣ Basic validation
	if user.Email == "" || user.Password == "" || user.Role == "" {
		http.Error(w, "Email, Password & Role required", http.StatusBadRequest)
		return
	}

	// 3️⃣ Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		http.Error(w, "Password error", 500)
		return
	}
	user.Password = string(hash)

	// 4️⃣ Save user
	userCollection := config.DB.Collection("users")

	// check already exists
	count, _ := userCollection.CountDocuments(context.TODO(), bson.M{"email": user.Email})
	if count > 0 {
		http.Error(w, "User already exists", 400)
		return
	}

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Register failed", 500)
		return
	}

	// 5️⃣ Response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"role": user.Role,
	})
}



// ================= LOGIN =================
func Login(w http.ResponseWriter, r *http.Request) {

	var input models.User
	var user models.User

	// 1️⃣ Decode input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}

	// 2️⃣ Find user
	collection := config.DB.Collection("users")

	err = collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", 404)
		return
	}

	// 3️⃣ Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Wrong password", 401)
		return
	}

	// 4️⃣ Generate JWT
	token, err := utils.GenerateJWT(user.ID.Hex(), user.Role)
	if err != nil {
		http.Error(w, "Token generation failed", 500)
		return
	}

	// 5️⃣ Send response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id": user.ID.Hex(),
			"email": user.Email,
			"role": user.Role,
		},
	})
}