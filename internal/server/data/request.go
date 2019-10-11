package data

import (
	"fmt"
	"net/http"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/gorilla/mux"
)

var ReadSafe func(string, HandlerFunction) error = Handlers.ReadSafe

func getStatus(res http.ResponseWriter, req *http.Request) {
	var method string
	var operation HandlerFunction = func(m *model.Handler) error {
		method = m.Request.Method
		return nil
	}

	vars := mux.Vars(req)
	hID := vars["handler_id"]
	err := ReadSafe(hID, operation)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = res.Write([]byte(method))
}

func getHost(res http.ResponseWriter, req *http.Request) {
	var host string
	var operation HandlerFunction = func(m *model.Handler) error {
		host = m.Request.Host
		return nil
	}

	vars := mux.Vars(req)
	hID := vars["handler_id"]
	err := ReadSafe(hID, operation)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = res.Write([]byte(host))
}
