package main

import (
	"log"
	"rami/database"
	"rami/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := database.InitDB(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	routes.GuestRoutes(router)
	routes.LogRoutes(router)
	routes.CSORoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
