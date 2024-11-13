	package routers

	import (
		"rami/controllers"

		"github.com/gorilla/mux"
	)

	func NewLogRouter(logController *controllers.LogController) *mux.Router {
		logRouter := mux.NewRouter()
		logRouter.HandleFunc("/logs/{serial}", logController.SearchLogsBySerialHandler).Methods("GET")
		logRouter.HandleFunc("/logs", logController.GetAllLogsHandler).Methods("GET")
		return logRouter
	}
