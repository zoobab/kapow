package data

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/gorilla/mux"
)

func TestSetStatus(t *testing.T) {
	request := httptest.NewRequest(http.MethodPut, "/handlers/HANDLER_XXXXXXXXXX/response/status", strings.NewReader("404"))
	response := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/handlers/{handler_id}/response/status", setStatus).Methods("PUT")

	handlerResponse := httptest.NewRecorder()
	myHandler := &model.Handler{
		ID:     "HANDLER_XXXXXXXXXX",
		Writer: handlerResponse,
	}

	WriteSafe = func(id string, f HandlerFunction) error {
		if id == myHandler.ID {
			return f(myHandler)
		}
		return errors.New("id not found")
	}

	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("HTTP Status mismatch. Expected: %d, got: %d", http.StatusOK, response.Code)
	}

	if handlerResponse.Code != 404 {
		t.Errorf("HTTP Status mismatch. Expected: %d, got: %d", http.StatusNotFound, handlerResponse.Code)
	}
}

func TestSetHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPut, "/handlers/HANDLER_XXXXXXXXXX/response/headers/pepe", strings.NewReader("mola"))
	response := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/handlers/{handler_id}/response/headers/{key}", setHeader).Methods("PUT")

	handlerResponse := httptest.NewRecorder()
	myHandler := &model.Handler{
		ID:     "HANDLER_XXXXXXXXXX",
		Writer: handlerResponse,
	}

	WriteSafe = func(id string, f HandlerFunction) error {
		if id == myHandler.ID {
			return f(myHandler)
		}
		return errors.New("id not found")
	}

	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("HTTP Status mismatch. Expected: %d, got: %d", http.StatusOK, response.Code)
	}

	headerValue := handlerResponse.Header().Get("pepe")
	if headerValue != "mola" {
		t.Errorf("Header value mismatch. Expected: %s, got: %s", "mola", headerValue)
	}
}
