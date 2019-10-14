package data

import (
	"net/http"

	"github.com/gorilla/mux"
)

var hasID func(string) bool = Handlers.Has

func Run(address string) error {
	r := configRouter()

	return http.ListenAndServe(address, r)
}

func configRouter() *mux.Router {
	r := mux.NewRouter()
	// the request tree
	r.HandleFunc("/handlers/{handler_id}/request/method", getStatus).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/host", getHost).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/path", getPath).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/matches/{key}", getMatches).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/params/{key}", getParams).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/headers/{key}", getHeader).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/cookies/{key}", getCookies).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/form/{key}", getForm).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/body", getBody).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/files/{file}/filename", getFileName).Methods("GET")
	r.HandleFunc("/handlers/{handler_id}/request/files/{file}/content", getFileContent).Methods("GET")

	// the response tree
	r.HandleFunc("/handlers/{handler_id}/response/status", setStatus).Methods("PUT")
	r.HandleFunc("/handlers/{handler_id}/response/headers/{key}", setHeader).Methods("PUT")
	r.HandleFunc("/handlers/{handler_id}/response/cookie/{key}", setCookie).Methods("PUT")
	r.HandleFunc("/handlers/{handler_id}/response/body", setBody).Methods("PUT")
	return r
}
