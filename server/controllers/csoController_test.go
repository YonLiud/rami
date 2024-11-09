package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rami/services"
	"rami/utils"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gorilla/mux"
)

var csoService *services.CSOService
var csoController *CSOController
var csoRouter *mux.Router

func initcsoControllerTest() {
	testDB := utils.InitiateTestDB()
	csoService = services.NewCSOService(testDB)
	csoController = NewCSOController(csoService)

	csoRouter = mux.NewRouter()
	csoRouter.HandleFunc("/cso/create", csoController.CreateCSOHandler).Methods("POST")
	csoRouter.HandleFunc("/cso/login", csoController.LoginCSOHandler).Methods("POST")
}

func TestCreatecsoHandler(t *testing.T) {
	initcsoControllerTest()

	csoReq := map[string]string{
		"username": gofakeit.Username(),
		"password": gofakeit.Password(true, true, true, true, false, 12),
	}

	body, _ := json.Marshal(csoReq)
	req, err := http.NewRequest("POST", "/cso/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	csoRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	cso, err := csoService.GetCSOByUsername(csoReq["username"])
	if err != nil {
		t.Fatal(err)
	}

	if cso.Username != csoReq["username"] {
		t.Errorf("handler returned unexpected username: got %v want %v", cso.Username, csoReq["username"])
	}
}

func TestCreatecsoHandlerInvalidRequest(t *testing.T) {
	initcsoControllerTest()

	csoReq := map[string]string{
		"username": "",
		"password": "",
	}

	body, _ := json.Marshal(csoReq)
	req, err := http.NewRequest("POST", "/cso/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	csoRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestCreatecsoHandlerUsernameExists(t *testing.T) {
	initcsoControllerTest()

	csoReq := map[string]string{
		"username": gofakeit.Username(),
		"password": gofakeit.Password(true, true, true, true, false, 12),
	}

	body, _ := json.Marshal(csoReq)
	req, err := http.NewRequest("POST", "/cso/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	csoRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	req, err = http.NewRequest("POST", "/cso/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	csoRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}
}

func TestLoginCSOHandler(t *testing.T) {
	initcsoControllerTest()

	csoReq := map[string]string{
		"username": gofakeit.Username(),
		"password": gofakeit.Password(true, true, true, true, false, 12),
	}

	body, _ := json.Marshal(csoReq)
	req, err := http.NewRequest("POST", "/cso/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	csoRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	req, err = http.NewRequest("POST", "/cso/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	csoRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	
}
