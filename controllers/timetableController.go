package controllers

import (
	"context"
	"net/http"
	"time"

	"school/config"
	"school/models"

	"github.com/gin-gonic/gin"
)

func AddTimetable(c *gin.Context) {
	var data models.Timetable
	c.BindJSON(&data)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config.DB.Collection("timetable").InsertOne(ctx, data)

	c.JSON(http.StatusOK, gin.H{"message": "Timetable Added"})
}