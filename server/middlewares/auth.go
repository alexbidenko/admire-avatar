package middlewares

import (
	"admire-avatar/config"
	"admire-avatar/entities"
	"admire-avatar/utils"
	"net/http"
)

type AuthorizedRequest struct {
	*http.Request
	User entities.BaseUser
}

type AuthorizedHandlerFunc func(w http.ResponseWriter, r *AuthorizedRequest)

func Auth(handler AuthorizedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		claims, err := utils.GetJWTClaims(tokenCookie.Value, []byte(config.AccessSecret))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		authorizedRequest := AuthorizedRequest{
			Request: r,
			User:    claims.BaseUser,
		}
		handler(w, &authorizedRequest)
	}
}
