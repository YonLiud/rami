package main

import (
	"log"
	"net/http"
	"rami/controllers"
	"rami/database"
	"rami/routers"
	"rami/services"

	"github.com/gorilla/mux"
)

const PORT string = ":3034"

func main() {

	log.Println("Initializing database...")
	db := database.InitDB("rami.db")

	log.Println("Starting services...")
	visitorService := services.NewVisitorService(db)
	logService := services.NewLogService(db)

	log.Println("Starting controllers...")
	visitorController := controllers.NewVisitorController(visitorService)
	logController := controllers.NewLogController(logService)

	log.Println("Starting routers...")
	visitorRouter := routers.NewVisitorRouter(visitorController)
	logRouter := routers.NewLogRouter(logController)

	log.Println("Starting main router...")
	mainRouter := mux.NewRouter()
	mainRouter.PathPrefix("/visitors").Handler(visitorRouter)
	mainRouter.PathPrefix("/logs").Handler(logRouter)

	log.Println("Starting Server on port ", PORT)
	http.ListenAndServe(PORT, mainRouter)

}
