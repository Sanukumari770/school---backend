package controllers

import (
	"context"
	"net/http"
	"time"

	"school/config"
	"school/models"

	"github.com/gin-gonic/gin"
)


// Subject
func AddSubject(c *gin.Context) {
	var data models.Subject
	c.BindJSON(&data)

	collection := config.DB.Collection("subjects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection.InsertOne(ctx, data)

	c.JSON(http.StatusOK, gin.H{"message": "Subject Added"})
}