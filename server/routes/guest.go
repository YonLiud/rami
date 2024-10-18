package routes

import (
	"rami/controllers"

	"github.com/gin-gonic/gin"
)

func GuestRoutes(r *gin.Engine) {
	guests := r.Group("/guests")
	{
		guests.POST("", controllers.CreateGuest)                 // Create a new guest
		guests.GET("", controllers.GetAllGuests)                 // Retrieve all
		guests.GET("/:id", controllers.GetGuestByID)             // Retrieve one
		guests.GET("/plate/:plate", controllers.GetGuestByPlate) // Retrieve by plate
		guests.PUT("/:id", controllers.UpdateGuest)              // Update a guest
		guests.DELETE("/:id", controllers.DeleteGuest)           // Delete a guest
		guests.POST("/enter/:id", controllers.MarkEntry)         // Mark entry for
		guests.POST("/exit/:id", controllers.MarkExit)           // Mark exit for
		guests.GET("/logs/:id", controllers.GetGuestLogs)        // Retrieve logs for
	}
}
