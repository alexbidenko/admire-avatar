package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"admire-avatar/ws"
	"archive/zip"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type GenerateRequestBody struct {
	utils.GeneratedImage
	Count int `json:"count"`
}

func GeneratePrints(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	var body GenerateRequestBody
	utils.ParseRequestBody(w, r.Request, &body)

	go func() {
		for i := 1; i <= body.Count; i++ {
			if i != 1 {
				time.Sleep(time.Second)
			}
			filename, err := utils.DownloadFile(body.GeneratedImage)
			if err != nil {
				fmt.Println(err)
			} else {
				image := entities.Image{
					BaseImage: entities.BaseImage{
						Source: filename,
					},
					UserID: r.User.ID,
					Type:   "print",
				}
				imageModule.Create(&image)

				err = os.Rename("files/temporary/"+image.Source, "files/images/"+image.Source)
				if err != nil {
					fmt.Println(err)
				}

				ws.PrintsPool.Broadcast <- ws.Message{Body: image, User: &r.User}
			}
		}
	}()

	utils.WriteJsonResponse(w, true)
}

func GetPaginatedPrints(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	offset, err := strconv.Atoi(mux.Vars(r.Request)["offset"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(mux.Vars(r.Request)["limit"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	images := imageModule.Paginate(r.User.ID, "print", offset, limit)

	utils.WriteJsonResponse(w, images)
}

func PrintToAvatar(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	imageID := mux.Vars(r.Request)["id"]

	err := imageModule.PrintToAvatar(r.User.ID, imageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteJsonResponse(w, true)
}

func Clear(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule

	images := imageModule.Get(r.User.ID, "print")
	for _, image := range images {
		err := os.Remove("files/images/" + image.Source)
		if err != nil {
			fmt.Println(err)
		}
	}

	imageModule.Clear(r.User.ID, "print")

	utils.WriteJsonResponse(w, true)
}

type DownloadArchiveBody struct {
}

func DownloadArchive(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	generateFileId := uuid.New().String()
	archive, err := os.Create("files/temporary/" + generateFileId + ".zip")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	var imagesModule modules.ImageModule
	images := imagesModule.Get(r.User.ID, "print")
	for _, image := range images {
		f1, err := os.Open("files/images/" + image.Source)
		if err == nil {
			defer f1.Close()

			w1, err := zipWriter.Create(image.Source)
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(w1, f1); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	zipWriter.Close()

	http.ServeFile(w, r.Request, "files/temporary/"+generateFileId+".zip")
}
