package controllers

import (
	"admire-avatar/config"
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"admire-avatar/ws"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetFolderImages(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]

	images := imageModule.GetByFolder(id, "avatar")

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

func GetImageFile(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
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

func GetImage(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var imageModule modules.ImageModule
	id := mux.Vars(r.Request)["id"]

	image, err := imageModule.Find(r.User.ID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteJsonResponse(w, image)
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

	defaultImage := func() {
		fileBytes, err := ioutil.ReadFile("default.jpg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		_, err = w.Write(fileBytes)
	}

	image, err := imageModule.GetAvatar(user.ID)
	if err != nil {
		defaultImage()
		return
	}

	fileBytes, err := ioutil.ReadFile("files/images/" + image.Source)
	if err != nil {
		defaultImage()
		return
	}

	_, err = w.Write(fileBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ShareRequestBody struct {
	UserId  uint `json:"user_id"`
	ImageId uint `json:"image_id"`
}

func ShareImage(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var data ShareRequestBody
	utils.ParseRequestBody(w, r.Request, &data)

	var folderModule modules.FolderModule
	var imageModule modules.ImageModule
	image, err := imageModule.Find(r.User.ID, strconv.Itoa(int(data.ImageId)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	folder, err := folderModule.GetByName(data.UserId, "Поделились с вами")
	if err != nil {
		folder = entities.Folder{
			UserID: data.UserId,
			Name:   "Поделились с вами",
		}
		folderModule.Create(&folder)
	}

	filename := uuid.New().String() + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"
	err = copy("files/images/"+image.Source, "files/images/"+filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	sharedImage := entities.Image{
		UserID:   data.UserId,
		FolderId: folder.ID,
		Type:     "avatar",
		Main:     false,
		BaseImage: entities.BaseImage{
			Source: filename,
		},
	}
	imageModule.Create(&sharedImage)

	var userModule modules.UserModule
	sharedUser, err := userModule.Find(data.UserId)
	if err == nil {
		go func() {
			ws.NotificationsPool.Broadcast <- ws.Message{Body: map[string]interface{}{
				"image": sharedImage,
				"user":  r.User,
				"type":  "share",
			}, User: &sharedUser.BaseUser}
		}()
	} else {
		fmt.Println(err)
	}

	utils.WriteJsonResponse(w, true)
}

func DownloadFolder(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	downloadFolder(w, r, mux.Vars(r.Request)["id"])
}

func GetTags(w http.ResponseWriter, _ *middlewares.AuthorizedRequest) {
	resp, err := http.Get(config.ServerApi + "/images/tags")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}
