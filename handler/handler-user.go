package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/shouryabansal7/BookFam/db"
	"github.com/shouryabansal7/BookFam/internal/database"
	"github.com/shouryabansal7/BookFam/models"
)

func HandlerCreateUser(apiCfg *db.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if apiCfg.DB == nil {
			http.Error(w, "Database not available", http.StatusInternalServerError)
			return
		}
		log.Println("Creating user with database connection...")
		type parameters struct {
			Name string
			Email string
			Password string
		}
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err!=nil {
			RespondWithError(w,400,fmt.Sprintf("Error parsing json",err))
			return
		}

		user, err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
			ID:        uuid.New(),
			Name:      params.Name,
			Email: 	   params.Email,
			Password:  params.Password,
		})

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create user: %v", err))
			return
		}
		RespondWithJSON(w,200,models.DatabaseUserToUser(user))
		// RespondWithJSON(w,200,struct{}{})

		fmt.Fprintln(w, "User created successfully!")
	}
}