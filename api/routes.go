package api

import (
	"github.com/aldaprojects/tinaap/tinaapp/web"
	"github.com/gorilla/mux"
)

func routes(userHTTPService *web.UserHTTPService) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/user", userHTTPService.PostUser).Methods("POST")
	r.HandleFunc("/user", userHTTPService.GetUser).Methods("GET")

	return r
}