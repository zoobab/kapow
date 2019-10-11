package data

import (
	"net/http"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/gorilla/mux"
)

var ReadSafe func(string, HandlerFunction) error = Handlers.ReadSafe

func performReadSafeOperation(res http.ResponseWriter, req *http.Request, operation HandlerFunction) error {
	vars := mux.Vars(req)
	hID := vars["handler_id"]
	return ReadSafe(hID, operation)
}

func getStatus(res http.ResponseWriter, req *http.Request) {
	var method string
	var operation HandlerFunction = func(m *model.Handler) error {
		method = m.Request.Method
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(method))
}

// TODO: check on real world where is the correct value
func getHost(res http.ResponseWriter, req *http.Request) {
	var host string
	var operation HandlerFunction = func(m *model.Handler) error {
		host = m.Request.Host
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(host))
}

func getPath(res http.ResponseWriter, req *http.Request) {
	var path string
	var operation HandlerFunction = func(m *model.Handler) error {
		path = m.Request.URL.EscapedPath()
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(path))
}

func getMatches(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	var value string
	var operation HandlerFunction = func(m *model.Handler) error {
		opVars := mux.Vars(m.Request)
		value = opVars[key]
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(value))
}

func getParams(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	var value string
	var operation HandlerFunction = func(m *model.Handler) error {
		value = m.Request.FormValue(key)
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(value))
}

func getHeader(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	var value string
	var operation HandlerFunction = func(m *model.Handler) error {
		value = m.Request.Header.Get(key)
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(value))
}

func getCookies(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	var value string
	var operation HandlerFunction = func(m *model.Handler) error {
		cookie, _ := m.Request.Cookie(key)
		if cookie.Name == key {
			value = cookie.Value
		}
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(value))
}

func getForm(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	var value string
	var operation HandlerFunction = func(m *model.Handler) error {
		value = m.Request.FormValue(key)
		return nil
	}

	_ = performReadSafeOperation(res, req, operation)

	_, _ = res.Write([]byte(value))
}
