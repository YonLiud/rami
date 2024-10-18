package routes

import (
	"rami/controllers"

	"github.com/gin-gonic/gin"
)

func LogRoutes(r *gin.Engine) {
	logs := r.Group("/logs")
	{
		logs.GET("", controllers.GetAllLogs)     // Retrieve all logs
		logs.GET("/:id", controllers.GetLogByID) // Retrieve log by guest ID
	}
}
