package data

import (
	"errors"
	"io"
	"net/http"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/gorilla/mux"
)

var ReadSafe func(string, HandlerFunction) error = Handlers.ReadSafe

// performReadSafeOperation only return error when ID not found
func performReadSafeOperation(res http.ResponseWriter, req *http.Request, operation HandlerFunction) error {
	vars := mux.Vars(req)
	hID := vars["handler_id"]
	has := hasID(hID)
	if !has {
		res.WriteHeader(http.StatusNotFound)
		return errors.New("Id Not Found")
	}
	return ReadSafe(hID, operation)
}

func getStatus(res http.ResponseWriter, req *http.Request) {
	var method string
	var operation HandlerFunction = func(m *model.Handler) error {
		method = m.Request.Method
		return nil
	}

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

	_, _ = res.Write([]byte(method))
}

// TODO: check on real world where is the correct value
func getHost(res http.ResponseWriter, req *http.Request) {
	var host string
	var operation HandlerFunction = func(m *model.Handler) error {
		host = m.Request.Host
		return nil
	}

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

	_, _ = res.Write([]byte(host))
}

func getPath(res http.ResponseWriter, req *http.Request) {
	var path string
	var operation HandlerFunction = func(m *model.Handler) error {
		path = m.Request.URL.EscapedPath()
		return nil
	}

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

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

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

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

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

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

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

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

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

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

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

	_, _ = res.Write([]byte(value))
}

func getBody(res http.ResponseWriter, req *http.Request) {
	var operation HandlerFunction = func(m *model.Handler) error {
		_, _ = io.Copy(res, m.Request.Body)
		return nil
	}

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}
}

func getFileName(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	theFile := vars["file"]
	var filename string
	var operation HandlerFunction = func(m *model.Handler) error {
		r := m.Request
		_ = r.ParseMultipartForm(10 << 20)
		_, handler, _ := r.FormFile(theFile)
		filename = handler.Filename
		return nil
	}

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}

	_, _ = res.Write([]byte(filename))
}

func getFileContent(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	theFile := vars["file"]
	var operation HandlerFunction = func(m *model.Handler) error {
		r := m.Request
		_ = r.ParseMultipartForm(10 << 20)
		file, _, _ := r.FormFile(theFile)

		_, _ = io.Copy(res, file)
		return nil
	}

	err := performReadSafeOperation(res, req, operation)
	if err != nil {
		return
	}
}
