package middleware

import (
	"fmt"
	"net/http"

	"github.com/shouryabansal7/BookFam/db"
	"github.com/shouryabansal7/BookFam/handler"
	"github.com/shouryabansal7/BookFam/internal/database"
	"github.com/shouryabansal7/BookFam/token"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func MiddlewareAuth(authedhandler authedHandler, apiCfg *db.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := token.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		if err != nil {
			handler.RespondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}
		claims, err := token.ExtractJWTInfo(tokenString)
		if err != nil {
			handler.RespondWithError(w, http.StatusNotFound, "Couldn't get claims")
			return
		}

		user, err := apiCfg.DB.GetUserByID(r.Context(), claims.UserID)
		if err != nil {
			handler.RespondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}
		//fmt.Sprintf("found it "+user.Email);
		authedhandler(w, r, user)
	}
}