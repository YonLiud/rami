package controllers

import (
	"net/http"
	"rami/database"
	"rami/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllLogs retrieves all logs from the database
func GetAllLogs(c *gin.Context) {
	var logs []models.Log

	if err := database.DB.Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

// GetLogByID retrieves a single log from the database using the log's ID number
func GetLogByID(c *gin.Context) {
	id := c.Param("id")
	var log models.Log

	if err := database.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}

func createLog(event string, guestID string) {
	log := models.Log{
		Event:     event,
		GuestID:   guestID,
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
	}

	if err := database.DB.Create(&log).Error; err != nil {
		panic(err)
	}
}
