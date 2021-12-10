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
	s.HandleFunc("/user/password", middlewares.Auth(controllers.ChangePassword)).Methods("POST")
	s.HandleFunc("/user", middlewares.Auth(controllers.GetUserByToken)).Methods("GET")
	s.HandleFunc("/user/change", middlewares.Auth(controllers.ChangeUser)).Methods("POST")

	return r
}
