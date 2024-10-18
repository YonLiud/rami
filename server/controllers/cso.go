package controllers

import (
	"net/http"
	"rami/database"
	"rami/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes the password using bcrypt
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// csoLogin Login for CSO
func CsoLogin(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cso models.CSO
	if err := database.DB.Where("username = ?", loginData.Username).First(&cso).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cso.HashedPassword), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.Set("isCSO", true)

	c.JSON(http.StatusOK, gin.H{"message": "CSO logged in successfully"})
}

// CreateCSO creates a new CSO in the database
func CreateCSO(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var cso models.CSO

	if err := c.ShouldBindJSON(&cso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cso.HashedPassword = hashPassword(password)

	if err := database.DB.Create(&cso).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createLog("CSO created", username)
	c.JSON(http.StatusCreated, cso)
}

// RemoveCSO deletes a CSO from the database
func RemoveCSO(c *gin.Context) {
	username := c.Param("username")

	var cso models.CSO
	if err := database.DB.Where("username = ?", username).First(&cso).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CSO not found"})
		return
	}

	if err := database.DB.Delete(&cso).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CSO removed successfully"})
}
