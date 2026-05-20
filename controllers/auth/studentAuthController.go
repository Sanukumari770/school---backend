package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
// for school models 
	"school/config"
	"school/models"
// for security and jwt token 
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
)

type StudentLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func StudentLoginController(w http.ResponseWriter, r *http.Request) {

	var loginData StudentLogin

	json.NewDecoder(r.Body).Decode(&loginData)

	collection := config.DB.Collection("students")

	var student models.Student

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"email": loginData.Email,
		},
	).Decode(&student)

	if err != nil {
		http.Error(w, "Student not found", http.StatusUnauthorized)
		return
	}

	// compare password
	println("DB HASH:", student.Password)
    println("USER PASSWORD:", loginData.Password)

err = bcrypt.CompareHashAndPassword(
	[]byte(student.Password),
	[]byte(loginData.Password),
)
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    student.ID.Hex(),
		"email": student.Email,
		"role":  "student",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		http.Error(w, "Token error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Student login successful",
		"token":   tokenString,
		"student": student,
	})
}