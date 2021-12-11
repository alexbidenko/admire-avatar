package main

import (
	"admire-avatar/controllers"
	"admire-avatar/middlewares"
	"admire-avatar/ws"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func initRoutes() http.Handler {
	go ws.PrintsPool.Start()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/user/sign-up", controllers.SignUp).Methods("POST")
	s.HandleFunc("/user/sign-in", controllers.SignIn).Methods("POST")
	s.HandleFunc("/user/logout", controllers.Logout).Methods("POST")
	s.HandleFunc("/user/refresh", controllers.Refresh).Methods("POST")
	s.HandleFunc("/user/password", middlewares.Auth(controllers.ChangePassword)).Methods("POST")
	s.HandleFunc("/user", middlewares.Auth(controllers.GetUserByToken)).Methods("GET")

	s.HandleFunc("/images", middlewares.Auth(controllers.SaveImage)).Methods("POST")
	s.HandleFunc("/images", middlewares.Auth(controllers.GenerateImage)).Methods("PUT")
	s.HandleFunc("/images/tags", middlewares.Auth(controllers.GetTags)).Methods("GET")
	s.HandleFunc("/images/avatar", middlewares.Auth(controllers.GetAvatar)).Methods("GET")
	s.HandleFunc("/images/folder/{id}", middlewares.Auth(controllers.GetFolderImages)).Methods("GET")
	s.HandleFunc("/images/{id}/folder/{folderId}", middlewares.Auth(controllers.ImageToFolder)).Methods("PUT")
	s.HandleFunc("/images/{id}", middlewares.Auth(controllers.RemoveImage)).Methods("DELETE")
	s.HandleFunc("/images/{id}", middlewares.Auth(controllers.CreateAvatar)).Methods("PUT")
	s.HandleFunc("/images/{id}", middlewares.Auth(controllers.GetImage)).Methods("GET")
	s.HandleFunc("/images/{offset}/{limit}", middlewares.Auth(controllers.GetPaginatedImages)).Methods("GET")

	s.HandleFunc("/prints/{offset}/{limit}", middlewares.Auth(controllers.GetPaginatedPrints)).Methods("GET")
	s.HandleFunc("/prints", middlewares.Auth(controllers.GeneratePrints)).Methods("POST")
	s.HandleFunc("/prints/{id}", middlewares.Auth(controllers.PrintToAvatar)).Methods("PUT")
	s.HandleFunc("/prints", middlewares.Auth(controllers.Clear)).Methods("DELETE")
	s.HandleFunc("/prints/archive", middlewares.Auth(controllers.DownloadArchive)).Methods("GET")

	s.HandleFunc("/folders", middlewares.Auth(controllers.GetFolders)).Methods("GET")
	s.HandleFunc("/folders", middlewares.Auth(controllers.CreateFolder)).Methods("POST")
	s.HandleFunc("/folders/{id}", middlewares.Auth(controllers.DeleteFolder)).Methods("DELETE")

	s.HandleFunc("/admire-avatar/{emailHash}", controllers.GetImageByEmail).Methods("GET")

	s.PathPrefix("/files/temporary/").Handler(http.StripPrefix("/api/files/temporary/", http.FileServer(http.Dir("files/temporary"))))
	s.PathPrefix("/files/images/").Handler(http.StripPrefix("/api/files/images/", http.FileServer(http.Dir("files/images"))))

	s.HandleFunc("/prints/channel", middlewares.Auth(func(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
		ws.ServeWs(ws.PrintsPool, w, r)
	}))

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
