package control

import (
	"net/http"
)

var (
	errorMalformedJSON = []byte(`{"reason": "Malformed JSON"}`)
	errorInvalidRoute  = []byte(`{"reason": "Invalid Route"}`)
	errorRouteNotFound = []byte(`{"reason": "Route Not Found"}`)
)

func writeError(w http.ResponseWriter, error []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, _ = w.Write(error)
}
