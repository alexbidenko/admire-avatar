package main

import (
	"admire-avatar/controllers"
	"admire-avatar/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func initRoutes() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/user/sign-up", controllers.SignUp).Methods("POST")
	s.HandleFunc("/user/sign-in", controllers.SignIn).Methods("POST")
	s.HandleFunc("/user/logout", controllers.Logout).Methods("POST")
	s.HandleFunc("/user/refresh", controllers.Refresh).Methods("POST")
	s.HandleFunc("/user/change-password", middlewares.Auth(controllers.ChangePassword)).Methods("POST")

	return r
}
