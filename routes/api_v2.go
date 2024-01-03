package routes

import (
	apiv2 "github.com/edalicioh/first_api_go/handlers/api/v2"
	"github.com/gorilla/mux"
)

func RegisterRoutesApiV2(router *mux.Router) {
	s := router.PathPrefix("/v2").Subrouter()

	s.HandleFunc("/", apiv2.HomeHandler).Methods("GET", "POST")
}
