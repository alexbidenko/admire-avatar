package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initRoutes() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/api").Subrouter()

	return r
}
