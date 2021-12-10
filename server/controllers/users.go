package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var userModule modules.UserModule
	var user entities.User
	utils.ParseRequestBody(w, r.Request, &user)

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := userModule.Find(r.User.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data.Password = string(bytes)
	userModule.Update(data.ID, &data)

	utils.WriteJsonResponse(w, nil)
}

func GetUserByToken(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	utils.WriteJsonResponse(w, r.User)
}

func ChangeUser(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var userModule modules.UserModule
	var changedUser entities.User
	utils.ParseRequestBody(w, r.Request, &changedUser)

	user, err := userModule.Find(r.User.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if changedUser.Name != "" {
		user.Name = changedUser.Name
	}
	if changedUser.Email != "" {
		user.Email = changedUser.Email
	}

	userModule.Update(user.ID, &user)

	utils.WriteJsonResponse(w, user)
}
