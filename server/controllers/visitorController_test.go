package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rami/models"
	"rami/services"
	"rami/utils"
	"testing"

	"github.com/gorilla/mux"
)

var visitorService *services.VisitorService
var visitorController *VisitorController
var visitorRouter *mux.Router

func initVisitorControllerTest() {
	testDB := utils.InitiateTestDB()
	visitorService = services.NewVisitorService(testDB)
	visitorController = NewVisitorController(visitorService)

	visitorRouter = mux.NewRouter()
	visitorRouter.HandleFunc("/visitors", visitorController.CreateVisitorHandler).Methods("POST")
	visitorRouter.HandleFunc("/visitors", visitorController.GetAllVisitorsHandler).Methods("GET")
	visitorRouter.HandleFunc("/visitors/{credentialsNumber}", visitorController.GetVisitorByCredentialsNumberHandler).Methods("GET")
	visitorRouter.HandleFunc("/visitors/{credentialsNumber}", visitorController.UpdateVisitorHandler).Methods("PUT")
	visitorRouter.HandleFunc("/visitors/{credentialsNumber}", visitorController.MarkEntryHandler).Methods("PATCH")
	visitorRouter.HandleFunc("/visitors/{credentialsNumber}", visitorController.MarkExitHandler).Methods("PATCH")
}

func TestCreateVisitorHandler(t *testing.T) {
	initVisitorControllerTest()

	visitorReq := utils.GenerateRandomVisitor()

	body, _ := json.Marshal(visitorReq)
	req, err := http.NewRequest("POST", "/visitors", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	visitorRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		t.FailNow()
	}
}

func TestCreateBadVisitorTest(t *testing.T) {
	initVisitorControllerTest()

	visitorReq := utils.GenerateRandomVisitor()
	visitorReq.CredentialsNumber = ""

	body, _ := json.Marshal(visitorReq)
	req, err := http.NewRequest("POST", "/visitors", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	visitorRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}
}

func TestCreateExistingVisitor(t *testing.T) {
	initVisitorControllerTest()

	visitorReq := utils.GenerateRandomVisitor()
	secondVisitorReq := utils.GenerateRandomVisitor()

	secondVisitorReq.CredentialsNumber = visitorReq.CredentialsNumber
	visitorService.CreateVisitor(&visitorReq)

	body, _ := json.Marshal(secondVisitorReq)
	req, err := http.NewRequest("POST", "/visitors", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	visitorRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGetAllVisitorsHandler(t *testing.T) {
	initVisitorControllerTest()

	visitors := make([]models.Visitor, 5)
	for i := 0; i < 5; i++ {
		visitor := utils.GenerateRandomVisitor()
		visitorService.CreateVisitor(&visitor)
		visitors[i] = visitor
	}

	req, err := http.NewRequest("GET", "/visitors", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	visitorRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var retrievedVisitors []models.Visitor
	if err := json.NewDecoder(rr.Body).Decode(&retrievedVisitors); err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if len(retrievedVisitors) != 5 {
		t.Errorf("Expected 5 visitors, got %d", len(retrievedVisitors))
	}

	for i := 0; i < 5; i++ {
		if visitors[i] != retrievedVisitors[i] {
			t.Errorf("Expected visitor %v, got %v", visitors[i], retrievedVisitors[i])
		}
	}
}

func TestGetVisitorByCredentialsNumberHandler(t *testing.T) {
	initVisitorControllerTest()

	visitor := utils.GenerateRandomVisitor()
	visitorService.CreateVisitor(&visitor)

	req, err := http.NewRequest("GET", "/visitors/"+visitor.CredentialsNumber, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	visitorRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var retrievedVisitor models.Visitor
	if err := json.NewDecoder(rr.Body).Decode(&retrievedVisitor); err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if visitor != retrievedVisitor {
		t.Errorf("Expected visitor %v, got %v", visitor, retrievedVisitor)
	}
}

func TestNonExistantGetVisitorByCredentialsNumberHandler(t *testing.T) {
	initVisitorControllerTest()

	req, err := http.NewRequest("GET", "/visitors/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	visitorRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

