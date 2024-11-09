package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"rami/models"
	"rami/services"
)

type CSOController struct {
	CSOService *services.CSOService
}

func NewCSOController(csoService *services.CSOService) *CSOController {
	return &CSOController{CSOService: csoService}
}

func (cc *CSOController) CreateCSOHandler(w http.ResponseWriter, r *http.Request) {
	type CSORequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var csoReq CSORequest
	if err := json.NewDecoder(r.Body).Decode(&csoReq); err != nil {
		http.Error(w, "Invalid request payload, "+err.Error(), http.StatusBadRequest)
		return
	}

	if csoReq.Username == "" || csoReq.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	err := cc.CSOService.CreateCSO(csoReq.Username, csoReq.Password)
	if err != nil {
		if errors.Is(err, models.ErrUsernameExists) {
			http.Error(w, "Username already in use", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to create CSO", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cc *CSOController) LoginCSOHandler(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request payload, "+err.Error(), http.StatusBadRequest)
		return
	}

	if loginReq.Username == "" || loginReq.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	_, err := cc.CSOService.AuthenticateCSO(loginReq.Username, loginReq.Password)

	if err != nil {
		if errors.Is(err, models.ErrCSONotFound) {
			http.Error(w, "CSO not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, models.ErrInvalidCredentials) {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Failed to authenticate CSO", http.StatusInternalServerError)
		return
	}

	//TODO generate token and return it

	w.WriteHeader(http.StatusOK)
}
