package data

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/gorilla/mux"
)

func TestGetMethod(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/handlers/HANDLER_XXXXXXXXXX/request/method", nil)
	response := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/handlers/{handler_id}/request/method", getStatus).Methods("GET")

	handlerRequest := httptest.NewRequest(http.MethodGet, "/foo/bar", nil)
	myHandler := &model.Handler{
		ID:      "HANDLER_XXXXXXXXXX",
		Request: handlerRequest,
	}

	ReadSafe = func(id string, f HandlerFunction) error {
		if id == myHandler.ID {
			return f(myHandler)
		}
		return errors.New("id not found")
	}

	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("HTTP Status mismatch. Expected: %d, got: %d", http.StatusOK, response.Code)
	}

	responseBytes, _ := ioutil.ReadAll(response.Body)
	retrieved := string(responseBytes)
	if retrieved != "GET" {
		t.Errorf("HTTP Method mismatch. Expected: %s, got: %s", "GET", retrieved)
	}
}

func TestGetHost(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/handlers/HANDLER_XXXXXXXXXX/request/host", nil)
	response := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/handlers/{handler_id}/request/host", getHost).Methods("GET")

	handlerRequest := httptest.NewRequest(http.MethodGet, "/foo/bar", nil)
	myHandler := &model.Handler{
		ID:      "HANDLER_XXXXXXXXXX",
		Request: handlerRequest,
	}

	ReadSafe = func(id string, f HandlerFunction) error {
		if id == myHandler.ID {
			return f(myHandler)
		}
		return errors.New("id not found")
	}

	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("HTTP Status mismatch. Expected: %d, got: %d", http.StatusOK, response.Code)
	}

	responseBytes, _ := ioutil.ReadAll(response.Body)
	retrieved := string(responseBytes)
	if retrieved != "example.com" {
		t.Errorf("Host mistmatch. Expected: %s, got: %s", "example.com", retrieved)
	}
}
