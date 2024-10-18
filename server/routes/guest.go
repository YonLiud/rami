package routes

import (
	"rami/controllers"
	"rami/middleware"

	"github.com/gin-gonic/gin"
)

func GuestRoutes(r *gin.Engine) {
	guests := r.Group("/guests")
	{
		guests.POST("", middleware.CheckCSO(), controllers.CreateGuest)       // Create a new guest (PROTECTED)
		guests.GET("", controllers.GetAllGuests)                              // Retrieve all
		guests.GET("/:id", controllers.GetGuestByID)                          // Retrieve one
		guests.GET("/plate/:plate", controllers.GetGuestByPlate)              // Retrieve by plate
		guests.PUT("/:id", middleware.CheckCSO(), controllers.UpdateGuest)    // Update a guest (PROTECTED)
		guests.DELETE("/:id", middleware.CheckCSO(), controllers.DeleteGuest) // Delete a guest (PROTECTED)
		guests.POST("/enter/:id", controllers.MarkEntry)                      // Mark entry for
		guests.POST("/exit/:id", controllers.MarkExit)                        // Mark exit for
		guests.GET("/logs/:id", controllers.GetGuestLogs)                     // Retrieve logs for
	}
}
