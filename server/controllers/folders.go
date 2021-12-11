package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPublic(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var folderModule modules.FolderModule

	folders := folderModule.GetPublic()

	utils.WriteJsonResponse(w, folders)
}

func GetFolders(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var folderModule modules.FolderModule

	folders := folderModule.Get(r.User.ID)

	utils.WriteJsonResponse(w, folders)
}

func GetFolder(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var folderModule modules.FolderModule
	id := mux.Vars(r.Request)["id"]

	folder, err := folderModule.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteJsonResponse(w, folder)
}

func CreateFolder(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var folderModule modules.FolderModule
	var folder entities.Folder
	utils.ParseRequestBody(w, r.Request, &folder)

	folder.UserID = r.User.ID
	folderModule.Create(&folder)

	utils.WriteJsonResponse(w, folder)
}

func DeleteFolder(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var folderModule modules.FolderModule
	id := mux.Vars(r.Request)["id"]

	folderModule.Delete(id)

	utils.WriteJsonResponse(w, true)
}

func UpdateFolder(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var folderModule modules.FolderModule
	var data entities.Folder
	utils.ParseRequestBody(w, r.Request, &data)
	id := mux.Vars(r.Request)["id"]

	folder, err := folderModule.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	folder.Name = data.Name
	folder.Public = data.Public
	folderModule.Update(r.User.ID, id, &folder)

	utils.WriteJsonResponse(w, folder)
}
