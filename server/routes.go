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

	s.HandleFunc("/images", middlewares.Auth(controllers.GetImages)).Methods("GET")
	s.HandleFunc("/images", middlewares.Auth(controllers.SaveImage)).Methods("POST")
	s.HandleFunc("/images", middlewares.Auth(controllers.GenerateImage)).Methods("PUT")
	s.HandleFunc("/images/{id}", middlewares.Auth(controllers.RemoveImage)).Methods("DELETE")
	s.HandleFunc("/images/{id}", middlewares.Auth(controllers.CreateAvatar)).Methods("PUT")
	s.HandleFunc("/images/{id}", middlewares.Auth(controllers.GetImage)).Methods("GET")

	s.HandleFunc("/admire-avatar/{emailHash}", controllers.GetImageByEmail).Methods("GET")

	s.PathPrefix("/files/temporary/").Handler(http.StripPrefix("/api/files/temporary/", http.FileServer(http.Dir("files/temporary"))))
	s.PathPrefix("/files/avatars/").Handler(http.StripPrefix("/api/files/avatars/", http.FileServer(http.Dir("files/avatars"))))

	return r
}
