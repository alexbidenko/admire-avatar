package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type GeneratedImage struct {
	Phrase string   `json:"phrase"`
	Tags   []string `json:"tags"`
}

func GetImages(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	images := imageModule.Get(r.User.ID)

	utils.WriteJsonResponse(w, images)
}

func CreateAvatar(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]

	imageModule.CreateAvatar(r.User.ID, id)

	utils.WriteJsonResponse(w, true)
}

func GenerateImage(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var data GeneratedImage
	utils.ParseRequestBody(w, r.Request, data)

	filename, err := downloadFile(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	image := &entities.BaseImage{
		Source: filename,
	}

	utils.WriteJsonResponse(w, image)
}

func SaveImage(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	var data entities.BaseImage
	utils.ParseRequestBody(w, r.Request, data)

	if strings.Contains(data.Source, "..") {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	err := os.Rename("images/temporary/"+data.Source, "image/avatars/"+data.Source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	image := entities.Image{
		BaseImage: data,
		UserId:    r.User.ID,
	}
	imageModule.Create(&image)

	utils.WriteJsonResponse(w, image)
}

func RemoveImage(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]

	image, err := imageModule.Find(r.User.ID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = os.Remove("image/avatars/" + image.Source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	imageModule.Delete(image.ID)

	utils.WriteJsonResponse(w, true)
}

func downloadFile(data GeneratedImage) (string, error) {
	body, err := json.Marshal(data)

	resp, err := http.Post("", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	filename := uuid.New().String() + "_" + time.Now().String() + ".png"
	out, err := os.Create("images/temporary/" + filename)
	if err != nil {
		return "", err
	}
	defer func() {
		err = out.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	return filename, err
}