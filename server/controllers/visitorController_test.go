package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
