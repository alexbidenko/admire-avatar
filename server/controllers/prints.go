package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
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
					UserId: r.User.ID,
					Type:   "print",
				}
				imageModule.Create(&image)

				err = os.Rename("files/temporary/"+image.Source, "files/images/"+image.Source)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	utils.WriteJsonResponse(w, true)
}

func GetPrints(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule

	images := imageModule.Get(r.User.ID, "print")

	utils.WriteJsonResponse(w, images)
}

func PrintToAvatar(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	imageID := mux.Vars(r.Request)["id"]

	imageModule.PrintToAvatar(r.User.ID, imageID)

	utils.WriteJsonResponse(w, true)
}
