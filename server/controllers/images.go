package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetFolderImages(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]

	images := imageModule.GetByFolder(r.User.ID, id, "avatar")

	utils.WriteJsonResponse(w, images)
}

func ImageToFolder(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]
	folderId := mux.Vars(r.Request)["folderId"]

	imageModule.ImageToFolder(r.User.ID, id, folderId)

	utils.WriteJsonResponse(w, true)
}

func GetPaginatedImages(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
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

	images := imageModule.Paginate(r.User.ID, "avatar", offset, limit)

	utils.WriteJsonResponse(w, images)
}

func CreateAvatar(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]

	imageModule.CreateAvatar(r.User.ID, id)

	utils.WriteJsonResponse(w, true)
}

func GenerateImage(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var data utils.GeneratedImage
	utils.ParseRequestBody(w, r.Request, &data)

	filename, err := utils.DownloadFile(data)
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
	utils.ParseRequestBody(w, r.Request, &data)

	if strings.Contains(data.Source, "..") {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	err := os.Rename("files/temporary/"+data.Source, "files/images/"+data.Source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	image := entities.Image{
		BaseImage: data,
		UserID:    r.User.ID,
		Type:      "avatar",
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

	err = os.Remove("files/images/" + image.Source)
	if err != nil {
		fmt.Println(err)
	}

	imageModule.Delete(image.ID)

	utils.WriteJsonResponse(w, true)
}

func GetImage(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]

	image, err := imageModule.Find(r.User.ID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fileBytes, err := ioutil.ReadFile("files/images/" + image.Source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write(fileBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetAvatar(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule

	avatar, err := imageModule.GetAvatar(r.User.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteJsonResponse(w, avatar)
}

func GetImageByEmail(w http.ResponseWriter, r *http.Request) {
	var userModule modules.UserModule
	var imageModule modules.ImageModule
	emailHash := mux.Vars(r)["emailHash"]

	user, err := userModule.GetByHash(emailHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	image, err := imageModule.GetAvatar(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fileBytes, err := ioutil.ReadFile("files/images/" + image.Source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write(fileBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTags(w http.ResponseWriter, _ *middlewares.AuthorizedRequest) {
	tags := []map[string]string{
		{
			"value": "abstractions",
			"label": "Абстракции",
		},
		{
			"value": "ar_deco",
			"label": "Арт-Деко",
		},
		{
			"value": "bosch",
			"label": "Иероним Босх",
		},
		{
			"value": "medieval",
			"label": "Средневековье",
		},
		{
			"value": "pop_art",
			"label": "Поп-Арт",
		},
		{
			"value": "van_goh",
			"label": "Ван Гог",
		},
	}
	utils.WriteJsonResponse(w, tags)
}
