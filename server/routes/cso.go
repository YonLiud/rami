package routes

import (
	"rami/controllers"
	"rami/middleware"

	"github.com/gin-gonic/gin"
)

func CSORoutes(r *gin.Engine) {
	cso := r.Group("/cso")
	{
		cso.POST("", middleware.CheckCSO(), controllers.CreateCSO)             // Create a new CSO 	(PROTECTED)
		cso.POST("/login", middleware.CheckCSO(), controllers.CsoLogin)        // Login for CSO 	(PROTECTED)
		cso.DELETE("/:username", middleware.CheckCSO(), controllers.RemoveCSO) // Delete a CSO 		(PROTECTED)
	}
}
