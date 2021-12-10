package middlewares

import (
	"admire-avatar/config"
	"admire-avatar/utils"
	"net/http"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {
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

		r.Header.Set("Email", claims.Email)
		handler.ServeHTTP(w, r)
	}
}
