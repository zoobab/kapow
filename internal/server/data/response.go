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

func performWriteSafeOperation(res http.ResponseWriter, req *http.Request, operation HandlerFunction) {
	vars := mux.Vars(req)

	hID := vars["handler_id"]
	has := hasID(hID)
	if !has {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	err := WriteSafe(hID, operation)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func setStatus(res http.ResponseWriter, req *http.Request) {
	value, _ := ioutil.ReadAll(req.Body)
	status, _ := strconv.Atoi(string(value))

	var operation HandlerFunction = func(m *model.Handler) error {
		m.Writer.WriteHeader(status)
		return nil
	}

	performWriteSafeOperation(res, req, operation)
}

func setHeader(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]

	value, _ := ioutil.ReadAll(req.Body)
	header := string(value)

	var operation HandlerFunction = func(m *model.Handler) error {
		m.Writer.Header().Set(key, header)
		return nil
	}

	performWriteSafeOperation(res, req, operation)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
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

	performWriteSafeOperation(res, req, operation)
}

func setBody(res http.ResponseWriter, req *http.Request) {
	var operation HandlerFunction = func(m *model.Handler) error {
		_, err := io.Copy(m.Writer, req.Body)
		return err
	}

	performWriteSafeOperation(res, req, operation)
}
