package data

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/gorilla/mux"
)

var WriteSafe func(string, HandlerFunction) error = Handlers.WriteSafe

func setStatus(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hID := vars["handler_id"]

	value, _ := ioutil.ReadAll(req.Body)
	status, _ := strconv.Atoi(string(value))

	var operation HandlerFunction = func(m *model.Handler) error {
		m.Writer.WriteHeader(status)
		return nil
	}

	err := WriteSafe(hID, operation)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func setHeader(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hID := vars["handler_id"]
	key := vars["key"]

	value, _ := ioutil.ReadAll(req.Body)
	header := string(value)

	var operation HandlerFunction = func(m *model.Handler) error {
		m.Writer.Header().Set(key, header)
		return nil
	}

	err := WriteSafe(hID, operation)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hID := vars["handler_id"]
	key := vars["key"]

	value, _ := ioutil.ReadAll(req.Body)
	cookieValue := string(value)

	cookie := http.Cookie{
		Name:  key,
		Value: cookieValue,
	}

	var operation HandlerFunction = func(m *model.Handler) error {
		http.SetCookie(m.Writer, &cookie)
		return nil
	}

	err := WriteSafe(hID, operation)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func setBody(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hID := vars["handler_id"]

	var operation HandlerFunction = func(m *model.Handler) error {
		io.Copy(m.Writer, req.Body)
		return nil
	}

	err := WriteSafe(hID, operation)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
