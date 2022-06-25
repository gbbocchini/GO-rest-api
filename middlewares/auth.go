package middlewares

import (
	"net/http"
	"rest/models"
	"rest/utils"
	"strings"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) < 2 {
			utils.SendError(w, http.StatusForbidden, models.Error{Message: "Token not provided!"})
			return
		}

		token := bearerToken[1]

		_, err := utils.VerifyJwtToken(token)
		if err != nil {
			utils.SendError(w, http.StatusForbidden, models.Error{Message: err.Error()})
			return
		}

		next.ServeHTTP(w, r)
	}
}
