package handler

import (
	"fmt"
	"net/http"

	"github.com/shouryabansal7/BookFam/db"
	"github.com/shouryabansal7/BookFam/internal/database"
)

func HandlerProtectedRoute(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	fmt.Fprint(w, user.Email);
	fmt.Fprint(w, "Welcome to the the protected area")
}