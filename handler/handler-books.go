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

func HandlerAddBook(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	if apiCfg.DB == nil {
		http.Error(w, "Database not available", http.StatusInternalServerError)
		return
	}
	log.Println("Creating user with database connection...")
	type parameters struct {
		Name string
		Author string
		Genre string
	}
	
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err!=nil {
		RespondWithError(w,400,fmt.Sprintf("Error parsing json: %v",err))
		return
	}

	book, err := apiCfg.DB.FindBookByName(r.Context(),params.Name)
	if err!=nil {
		fmt.Fprintln(w,"Book does not already exist in db")
		fmt.Fprintln(w,"Create Book in db")

		book, err :=apiCfg.DB.CreateBook(r.Context(),database.CreateBookParams{
			ID:        uuid.New(),
			Name:      params.Name,
			Author: 	   params.Author,
			Genre:  params.Genre,
		})

		if err != nil {
			log.Println(err)
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create user: %v", err))
			return
		}

		updateParams := database.AddUserToBookParams{
			ID:     book.ID,
			ArrayAppend: user.ID,
		}

		err = apiCfg.DB.AddUserToBook(r.Context(), updateParams)
		if err != nil {
			log.Printf("Error updating user_ids: %v", err)
			RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to add user to the book: %v", err))
			return
		}

		addBookIdToUserList(w,r,book,user,apiCfg)

		RespondWithJSON(w,200,fmt.Sprintf("User successfully added to book '%s'", params.Name))
	}else{
		for _, id := range book.UserIds {
			if id == user.ID {
				// If the user_id is already in the array, we don't need to add it
				RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("User already added: %v", err))
				return
			}
		}

		updateParams := database.AddUserToBookParams{
			ID:     book.ID,
			ArrayAppend: user.ID,
		}

		err = apiCfg.DB.AddUserToBook(r.Context(), updateParams)
		if err != nil {
			log.Printf("Error updating user_ids: %v", err)
			RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to add user to the book: %v", err))
			return
		}

		addBookIdToUserList(w,r,book,user,apiCfg)

		// Step 4: Respond with a success message
		RespondWithJSON(w,200,fmt.Sprintf("User successfully added to book '%s'", params.Name))
	}
}

func addBookIdToUserList(w http.ResponseWriter, r *http.Request,book database.Book, user database.User, apiCfg *db.ApiConfig){
	updateParams := database.AddBookIDToUserParams{
		ID:     user.ID,
		ArrayAppend: book.ID,
	}
	err := apiCfg.DB.AddBookIDToUser(r.Context(), updateParams)
	if err != nil {
		log.Printf("Error updating user db: %v", err)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to add book to the user's array: %v", err))
		return
	}

	RespondWithJSON(w,200,fmt.Sprintf("Book successfully added to user array '%s'", user.Name))
}

func HandlerGetBooks(w http.ResponseWriter, r *http.Request, user database.User, apiCfg *db.ApiConfig) {
	books, err := apiCfg.DB.GetBooks(r.Context())

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get books: %v",err))
		return
	}


	RespondWithJSON(w,200,models.DatabaseBooksToBooks(books))
}