package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func GetFolders(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var folderModule modules.FolderModule

	folders := folderModule.Get(r.User.ID)

	utils.WriteJsonResponse(w, folders)
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
