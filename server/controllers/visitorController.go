package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rami/models"
	"rami/services"

	"github.com/gorilla/mux"
)

type VisitorController struct {
	VisitorService *services.VisitorService
}

func NewVisitorController(visitorService *services.VisitorService) *VisitorController {
	return &VisitorController{VisitorService: visitorService}
}

func (vc *VisitorController) GetAllVisitorsHandler(w http.ResponseWriter, r *http.Request) {
	visitors, err := vc.VisitorService.GetAllVisitors()
	if err != nil {
		log.Printf("Failed to retrieve visitors: %v", err)
		http.Error(w, "Failed to retrieve visitors", http.StatusInternalServerError)
		return
	}

	if len(visitors) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(visitors)
}

func (vc *VisitorController) CreateVisitorHandler(w http.ResponseWriter, r *http.Request) {
	var visitor models.Visitor

	if err := json.NewDecoder(r.Body).Decode(&visitor); err != nil {
		http.Error(w, "Invalid request payload, "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := vc.VisitorService.CreateVisitor(&visitor); err != nil {
		log.Printf("Failed to create visitor: %v", err)

		if err.Error() == "UNIQUE constraint failed: visitors.credentials_number" {
			http.Error(w, "Credentials number is already in use", http.StatusBadRequest)
			return
		}

		if err.Error() == "NOT NULL constraint failed: visitors.credentials_number" {
			http.Error(w, "Credentials number is required", http.StatusUnprocessableEntity)
		}

		http.Error(w, "Failed to create visitor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (vc *VisitorController) GetVisitorByCredentialsNumberHandler(w http.ResponseWriter, r *http.Request) {
	credentialsNumber := mux.Vars(r)["credentialsNumber"]

	if credentialsNumber == "" {
		http.Error(w, "Credentials number is required", http.StatusBadRequest)
		return
	}

	visitor, err := vc.VisitorService.GetVisitorByCredentialsNumber(credentialsNumber)
	if err != nil {
		log.Printf("Failed to retrieve visitor: %v", err)
		http.Error(w, "Failed to retrieve visitor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(visitor)
}

func (vc *VisitorController) UpdateVisitorHandler(w http.ResponseWriter, r *http.Request) {
	credentialsNumber := mux.Vars(r)["credentialsNumber"]

	if credentialsNumber == "" {
		http.Error(w, "Credentials number is required", http.StatusBadRequest)
		return
	}

	var visitor models.Visitor
	if err := json.NewDecoder(r.Body).Decode(&visitor); err != nil {
		http.Error(w, "Invalid request payload, "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := vc.VisitorService.UpdateVisitor(&visitor); err != nil {
		log.Printf("Failed to update visitor: %v", err)
		http.Error(w, "Failed to update visitor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (vc *VisitorController) MarkEntryHandler(w http.ResponseWriter, r *http.Request) {
	credentialsNumber := mux.Vars(r)["credentialsNumber"]

	if credentialsNumber == "" {
		http.Error(w, "Credentials number is required", http.StatusBadRequest)
		return
	}

	visitor, err := vc.VisitorService.GetVisitorByCredentialsNumber(credentialsNumber)
	if err != nil {
		log.Printf("Failed to retrieve visitor: %v", err)
		http.Error(w, "Failed to retrieve visitor", http.StatusInternalServerError)
		return
	}

	if visitor.IsInside {
		http.Error(w, "Visitor is already inside", http.StatusBadRequest)
		return
	}

	if err := vc.VisitorService.MarkEntry(&visitor); err != nil {
		log.Printf("Failed to mark entry: %v", err)
		http.Error(w, "Failed to mark entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (vc *VisitorController) MarkExitHandler(w http.ResponseWriter, r *http.Request) {
	credentialsNumber := mux.Vars(r)["credentialsNumber"]

	if credentialsNumber == "" {
		http.Error(w, "Credentials number is required", http.StatusBadRequest)
		return
	}

	visitor, err := vc.VisitorService.GetVisitorByCredentialsNumber(credentialsNumber)
	if err != nil {
		log.Printf("Failed to retrieve visitor: %v", err)
		http.Error(w, "Failed to retrieve visitor", http.StatusInternalServerError)
		return
	}

	if !visitor.IsInside {
		http.Error(w, "Visitor is already outside", http.StatusBadRequest)
		return
	}

	if err := vc.VisitorService.MarkExit(&visitor); err != nil {
		log.Printf("Failed to mark exit: %v", err)
		http.Error(w, "Failed to mark exit", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
