package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rami/models"
	"rami/services"

	"github.com/gorilla/mux"
)

type LogController struct {
	LogService *services.LogService
}

func NewLogController(logService *services.LogService) *LogController {
	return &LogController{LogService: logService}
}

func (lc *LogController) CreateLogHandler(w http.ResponseWriter, r *http.Request) {
	var logEntry models.Log

	type LogRequest struct {
		Serial string `json:"serial"`
		Event  string `json:"event"`
	}

	var logReq LogRequest
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		http.Error(w, "Invalid request payload, "+err.Error(), http.StatusBadRequest)
		return
	}

	logEntry.Serial = logReq.Serial
	logEntry.Event = logReq.Event
	logEntry.Timestamp = models.GetCurrentTimestamp()
	if err := lc.LogService.CreateLog(&logEntry); err != nil {
		log.Printf("Failed to create log: %v", err)
		http.Error(w, "Failed to create log", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (lc *LogController) SearchLogsBySerialHandler(w http.ResponseWriter, r *http.Request) {
	serial := mux.Vars(r)["serial"]

	if serial == "" {
		http.Error(w, "Serial is required", http.StatusBadRequest)
		return
	}

	logs, err := lc.LogService.GetLogsBySerial(serial)
	if err != nil {
		log.Printf("Failed to retrieve logs: %v", err)
		http.Error(w, "Failed to retrieve logs", http.StatusInternalServerError)
		return
	}

	if len(logs) == 0 {
		http.Error(w, "No logs found", http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func (lc *LogController) GetAllLogsHandler(w http.ResponseWriter, r *http.Request) {
	logs, err := lc.LogService.GetAllLogs()
	if err != nil {
		log.Printf("Failed to retrieve logs: %v", err)
		http.Error(w, "Failed to retrieve logs", http.StatusInternalServerError)
		return
	}

	if len(logs) == 0 {
		http.Error(w, "No logs found", http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
