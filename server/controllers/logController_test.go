package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rami/models"
	"rami/services"
	"rami/utils"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gorilla/mux"
)

var logService *services.LogService
var logController *LogController
var logRouter *mux.Router

func initLogControllerTest() {
	testDB := utils.InitiateTestDB()
	logService = services.NewLogService(testDB)
	logController = NewLogController(logService)

	logRouter = mux.NewRouter()
	logRouter.HandleFunc("/logs/{serial}", logController.SearchLogsBySerialHandler).Methods("GET")
	logRouter.HandleFunc("/logs", logController.GetAllLogsHandler).Methods("GET")
}

func TestSearchLogsBySerialHandler(t *testing.T) {
	initLogControllerTest()

	serial := gofakeit.UUID()
	err := logService.CreateLog(&models.Log{
		Serial: serial,
		Event:  gofakeit.Sentence(3),
	})
	if err != nil {
		t.Fatalf("Error creating log: %v", err)
	}

	req, err := http.NewRequest("GET", "/logs/"+serial, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	rr := httptest.NewRecorder()
	logRouter.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var retrievedLogs []models.Log
	if err := json.NewDecoder(rr.Body).Decode(&retrievedLogs); err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if len(retrievedLogs) != 1 {
		t.Errorf("Expected 1 log, got %d", len(retrievedLogs))
	}
}

func TestGetAllLogsHandler(t *testing.T) {
	initLogControllerTest()

	req, err := http.NewRequest("GET", "/logs", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	logRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
		t.FailNow()
	}

	randLogs := make([]models.Log, 5)
	for i := range randLogs {
		randLogs[i] = models.Log{
			Serial: gofakeit.UUID(),
			Event:  gofakeit.Sentence(3),
		}
		if err := logService.CreateLog(&randLogs[i]); err != nil {
			t.Fatalf("Error creating log: %v", err)
		}
	}

	req, err = http.NewRequest("GET", "/logs", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	logRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var logs []models.Log
	if err := json.NewDecoder(rr.Body).Decode(&logs); err != nil {
		t.Fatal(err)
	}
	if len(logs) == 0 {
		t.Errorf("expected logs to be returned, got none")
	}
}
