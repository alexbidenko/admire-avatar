package controllers

import (
	"admire-avatar/entities"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var userModule modules.UserModule
	var user entities.User
	utils.ParseRequestBody(w, r, &user)

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := userModule.GetByEmail(r.Header.Get("Email"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data.Password = string(bytes)
	userModule.Update(data.ID, &data)

	utils.WriteJsonResponse(w, nil)
}

func GetUserByToken(w http.ResponseWriter, r *http.Request) {
	var userModule modules.UserModule
	user, err := userModule.GetByEmail(r.Header.Get("Email"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	utils.WriteJsonResponse(w, user.BaseUser)
}

func ChangeUser(w http.ResponseWriter, r *http.Request) {
	var userModule modules.UserModule
	var changedUser entities.User
	utils.ParseRequestBody(w, r, &changedUser)

	user, err := userModule.Find(r.Header.Get("Id"))
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
