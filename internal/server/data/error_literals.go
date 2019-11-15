package data

import (
	"net/http"
)

var (
	errorHandlerIDNotFound    = []byte(`{"reason": "Handler ID Not Found"}`)
	errorInvalidResourcePath  = []byte(`{"reason": "Invalid Resource Path"}`)
	errorResourceItemNotFound = []byte(`{"reason": "Resource Item Not Found"}`)
)

func writeError(w http.ResponseWriter, error []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, _ = w.Write(error)
}
