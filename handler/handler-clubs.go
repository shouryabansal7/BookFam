package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/shouryabansal7/BookFam/db"
	"github.com/shouryabansal7/BookFam/internal/database"
	"github.com/shouryabansal7/BookFam/models"
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

func HandlerJoinClub(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	clubJoinIDstr := chi.URLParam(r, "ClubID")
	clubJoinID, err := uuid.Parse(clubJoinIDstr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid club join ID")
		return
	}

	updateParams := database.AddUserClubParams{
		ID:     clubJoinID,
		ArrayAppend: user.ID,
	}

	err = apiCfg.DB.AddUserClub(r.Context(), updateParams)
	if err != nil {
		log.Printf("Error updating user_ids: %v", err)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to add user to the club: %v", err))
		return
	}

	RespondWithJSON(w,200,fmt.Sprintf("User successfully added to club"))
}

func HandlerLeaveClub(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	clubJoinIDstr := chi.URLParam(r, "ClubID")
	clubJoinID, err := uuid.Parse(clubJoinIDstr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid club leave ID")
		return
	}

	updateParams := database.RemoveUserFromClubParams{
		ID:     clubJoinID,
		ArrayRemove: user.ID,
	}

	err = apiCfg.DB.RemoveUserFromClub(r.Context(), updateParams)
	if err != nil {
		log.Printf("Error updating user_ids: %v", err)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to remove user from the club: %v", err))
		return
	}

	RespondWithJSON(w,200,fmt.Sprintf("User successfully left to club"))
}

func HandlerGetClubs(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	clubs, err := apiCfg.DB.GetClubs(r.Context())

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get clubs: %v",err))
		return
	}


	RespondWithJSON(w,200,models.DatabaseClubsToClubs(clubs))
}

func HandlerClubMembers(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	clubIDstr := chi.URLParam(r, "ClubID")
	clubID, err := uuid.Parse(clubIDstr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid club join ID")
		return
	}

	

	club, err := apiCfg.DB.GetClubByID(r.Context(),clubID)
	if err != nil {
		log.Printf("Error finding the club: %v", err)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to fetch members of the club with club id: %v", err))
		return
	}

	result := make([]string, len(club.UserIds))

	for i, userId := range club.UserIds{
		user, err := apiCfg.DB.GetUserByID(r.Context(),userId)
		if err != nil {
			log.Printf("Error finding the user details: %v", err)
			RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to fetch members of the club with club id: %v", err))
			return
		}
		result[i] = user.Name
	}

	RespondWithJSON(w,200,result)
}