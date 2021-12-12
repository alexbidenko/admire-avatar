package controllers

import (
	"admire-avatar/config"
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"admire-avatar/modules"
	"admire-avatar/utils"
	"context"
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var userModule modules.UserModule
	var user entities.User
	utils.ParseRequestBody(w, r, &user)

	_, err := userModule.GetByEmail(user.Email)
	if err == nil {
		http.Error(w, "Email already in use", http.StatusBadRequest)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = string(bytes)
	user.EmailHash = fmt.Sprintf("%x", md5.Sum([]byte(user.Email)))

	userModule.Create(&user)

	writeTokens(w, &user)
	utils.WriteJsonResponse(w, user.BaseUser)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var userModule modules.UserModule

	var authentication Authentication
	utils.ParseRequestBody(w, r, &authentication)

	user, err := userModule.GetByEmail(authentication.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authentication.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	writeTokens(w, &user)
	utils.WriteJsonResponse(w, user.BaseUser)
}

func ChangePassword(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var userModule modules.UserModule
	var user entities.User
	utils.ParseRequestBody(w, r.Request, &user)

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := userModule.GetByEmail(r.User.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data.Password = string(bytes)
	userModule.Update(data.ID, &data)

	utils.WriteJsonResponse(w, nil)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		removeTokens(w)
		http.Error(w, "Refresh token is required", http.StatusUnauthorized)
		return
	}

	claims, err := utils.GetJWTClaims(refreshTokenCookie.Value, []byte(config.RefreshSecret))
	if err != nil {
		removeTokens(w)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	_, err = config.RedisClient.Get(ctx, claims.Uuid).Result()
	if err != nil {
		removeTokens(w)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	var userModule modules.UserModule
	user, err := userModule.GetByEmail(claims.Email)
	if err != nil {
		removeTokens(w)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeTokens(w, &user)
	config.RedisClient.Del(ctx, claims.Uuid)

	utils.WriteJsonResponse(w, true)
}

func SearchUsers(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	email := r.URL.Query().Get("email")
	var userModule modules.UserModule

	users := userModule.Search(email, r.User.ID)

	utils.WriteJsonResponse(w, users)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	accessTokenCookie, err := r.Cookie("access_token")
	if err == nil {
		claims, err := utils.GetJWTClaims(accessTokenCookie.Value, []byte(config.AccessSecret))
		if err == nil {
			config.RedisClient.Del(ctx, claims.Uuid)
		}
	}
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err == nil {
		claims, err := utils.GetJWTClaims(refreshTokenCookie.Value, []byte(config.RefreshSecret))
		if err == nil {
			config.RedisClient.Del(ctx, claims.Uuid)
		}
	}
	removeTokens(w)

	utils.WriteJsonResponse(w, true)
}

func removeTokens(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "refresh_token",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "authorized",
		Path:   "/",
		MaxAge: -1,
	})
}

func writeTokens(w http.ResponseWriter, user *entities.User) {
	accessToken, err := user.GenerateAccessToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := user.GenerateRefreshToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		Secure:   os.Getenv("MODE") != "production",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Secure:   os.Getenv("MODE") != "production",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "authorized",
		Value:   "true",
		Path:    "/",
		Secure:  os.Getenv("ADMIN_HTTPS") != "false",
		Expires: time.Now().Add(time.Hour * 24 * 30),
	})
}

func GetUserByToken(w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	var userModule modules.UserModule

	user, err := userModule.Find(r.User.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	utils.WriteJsonResponse(w, user)
}
