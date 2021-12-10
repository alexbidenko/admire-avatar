package main

import (
	"admire-avatar/controllers"
	"admire-avatar/middlewares"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	if os.Getenv("MODE") == "production" {
		r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join("dist", r.URL.Path)

			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				http.ServeFile(w, r, filepath.Join("dist", "index.html"))
				return
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if strings.HasSuffix(r.URL.Path, ".js") {
				w.Header().Set("Content-Type", "application/javascript")
			}
			http.FileServer(http.Dir("dist")).ServeHTTP(w, r)
		})
	}

	return r
}
