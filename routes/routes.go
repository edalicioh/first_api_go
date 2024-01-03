package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutesApi(router *mux.Router) {
	s := router.PathPrefix("/api").Subrouter()
	RegisterRoutesApiV1(s)
	RegisterRoutesApiV2(s)
}
