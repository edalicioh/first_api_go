package routes

import (
	apiv1 "github.com/edalicioh/first_api_go/handlers/api/v1"
	"github.com/edalicioh/first_api_go/middlewares"
	"github.com/gorilla/mux"
)

func RegisterRoutesApiV1(router *mux.Router) {
	s := router.PathPrefix("/v1").Subrouter()

	s.HandleFunc("/", apiv1.HomeHandler).Methods("GET")

	authRoute := s.PathPrefix("/auth").Subrouter()

	authRoute.HandleFunc("/user", apiv1.UserHandler).Methods("GET")
	authRoute.Use(middlewares.AuthMiddlerwares)
}
