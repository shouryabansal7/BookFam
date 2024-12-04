package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/shouryabansal7/BookFam/db"
	"github.com/shouryabansal7/BookFam/internal/database"
)

func HandlerCreateClub(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	if apiCfg.DB == nil {
		http.Error(w, "Database not available", http.StatusInternalServerError)
		return
	}
	log.Println("Creating club with database connection...")
	type parameters struct {
		Name string
		Genre string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err!=nil {
		RespondWithError(w,400,fmt.Sprintf("Error parsing json: %v",err))
		return
	}

	club, err := apiCfg.DB.CreateClubs(r.Context(),database.CreateClubsParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Genre:  params.Genre,
	})

	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create club: %v", err))
		return
	}

	updateParams := database.AddUserClubParams{
		ID:     club.ID,
		ArrayAppend: user.ID,
	}

	err = apiCfg.DB.AddUserClub(r.Context(), updateParams)
	if err != nil {
		log.Printf("Error updating user_ids: %v", err)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to add user to the book: %v", err))
		return
	}

	RespondWithJSON(w,200,fmt.Sprintf("User successfully added to club '%s'", params.Name))
}