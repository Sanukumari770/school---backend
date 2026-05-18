package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"school/config"
	"school/models"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type TeacherLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TeacherLoginController(w http.ResponseWriter, r *http.Request) {

	var loginData TeacherLogin

	json.NewDecoder(r.Body).Decode(&loginData)

	collection := config.DB.Collection("teachers")

	var teacher models.Teacher

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"email": loginData.Email,
		},
	).Decode(&teacher)

	if err != nil {
		http.Error(w, "Teacher not found", http.StatusUnauthorized)
		return
	}

	inputPassword := strings.TrimSpace(loginData.Password)

	println("DB HASH:", teacher.Password)
	println("USER PASSWORD:", inputPassword)

	err = bcrypt.CompareHashAndPassword(
		[]byte(teacher.Password),
		[]byte(inputPassword),
	)

	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// JWT TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    teacher.ID.Hex(),
		"email": teacher.Email,
		"role":  "teacher",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		http.Error(w, "Token error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Teacher login successful",
		"token":   tokenString,
		"teacher": teacher,
	})
}
