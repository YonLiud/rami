package controllers

import (
	"fmt"
	"net/http"
	"rami/database"
	"rami/models"

	"github.com/gin-gonic/gin"
)

// GetAllGuests retrieves all guests from the database
func GetAllGuests(c *gin.Context) {
	var guests []models.Guest

	if err := database.DB.Find(&guests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guests)
}

// GetGuestByID retrieves a single guest from the database using the guest's ID number
func GetGuestByID(c *gin.Context) {
	idNumber := c.Param("id")
	var guest models.Guest

	if err := database.DB.Where("id_number = ?", idNumber).First(&guest).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}

	c.JSON(http.StatusOK, guest)
}

// GetGuestByPlate retrieves all guests from the database using the provided vehicle plate
func GetGuestByPlate(c *gin.Context) {
	plate := c.Param("plate")
	var guests []models.Guest

	if err := database.DB.Where("vehicle_plate = ?", plate).Find(&guests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guests)
}

// CreateGuest creates a new guest in the database
func CreateGuest(c *gin.Context) {
	var guest models.Guest

	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	createLog("created", fmt.Sprintf("%d", guest.ID))
	c.JSON(http.StatusCreated, guest)
}

// UpdateGuest updates an existing guest in the database
func UpdateGuest(c *gin.Context) {
	id := c.Param("id")
	var guest models.Guest

	if err := database.DB.First(&guest, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}

	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createLog("updated", fmt.Sprintf("%d", guest.ID))
	c.JSON(http.StatusOK, guest)
}

// DeleteGuest deletes a guest from the database
func DeleteGuest(c *gin.Context) {
	id := c.Param("id")
	var guest models.Guest

	if err := database.DB.First(&guest, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}

	if err := database.DB.Delete(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createLog("deleted", fmt.Sprintf("%d", guest.ID))

	c.Status(http.StatusNoContent)
}

// GetGuestLogs retrieves all logs associated with a specific guest
func GetGuestLogs(c *gin.Context) {
	id := c.Param("id")
	var logs []models.Log

	if err := database.DB.Where("guest_id = ?", id).Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

// MarkEntry marks a guest as inside the premises
func MarkEntry(c *gin.Context) {
	id := c.Param("id")
	var guest models.Guest

	if err := database.DB.First(&guest, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}

	if guest.IsInside {
		c.JSON(http.StatusUnavailableForLegalReasons, gin.H{"error": "Guest is already inside"})
		return
	}

	if err := database.DB.Save(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createLog("entry", fmt.Sprintf("%d", guest.ID))

	c.JSON(http.StatusOK, guest)
}

// MarkExit marks a guest as outside the premises
func MarkExit(c *gin.Context) {
	id := c.Param("id")
	var guest models.Guest

	if err := database.DB.First(&guest, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}

	if !guest.IsInside {
		c.JSON(http.StatusUnavailableForLegalReasons, gin.H{"error": "Guest is already outside"})
		return
	}
	guest.IsInside = false

	if err := database.DB.Save(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createLog("exit", fmt.Sprintf("%d", guest.ID))
	c.JSON(http.StatusOK, guest)
}
