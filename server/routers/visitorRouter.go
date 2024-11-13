package routers

import (
	"rami/controllers"

	"github.com/gorilla/mux"
)

func NewVisitorRouter(visitorController *controllers.VisitorController) *mux.Router {
	visitorRouter := mux.NewRouter()
	visitorRouter.HandleFunc("/visitors", visitorController.GetAllVisitorsHandler).Methods("GET")
	visitorRouter.HandleFunc("/visitors", visitorController.CreateVisitorHandler).Methods("POST")
	visitorRouter.HandleFunc("/visitors/inside", visitorController.GetAllVisitorsInsideHandler).Methods("GET")
	visitorRouter.HandleFunc("/visitors/{credentialsNumber}", visitorController.GetVisitorByCredentialsNumberHandler).Methods("GET")
	visitorRouter.HandleFunc("/visitors/{credentialsNumber}", visitorController.UpdateVisitorHandler).Methods("PUT")
	visitorRouter.HandleFunc("/visitors/{credentialsNumber}", visitorController.MarkEntryExitHandler).Methods("PATCH")
	return visitorRouter
}
