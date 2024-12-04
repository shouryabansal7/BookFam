package models

import (
	"github.com/google/uuid"
	"github.com/shouryabansal7/BookFam/internal/database"
)
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string 	`json:"email"`
	Password  string	`json:"password"`
}

func DatabaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email: 	   user.Email,
		Password:  user.Password,	
	}
}

type Book struct {
	ID        	uuid.UUID 	`json:"id"`
	Name      	string   	`json:"name"`
	Author     	string 		`json:"author"`
	Genre  		string		`json:"genre"`
}

func DatabaseBookToBook(book database.Book) Book {
	return Book{
		ID:        	book.ID,
		Name:      	book.Name,
		Author: 	book.Author,
		Genre:  	book.Genre,	
	}
}

func DatabaseBooksToBooks(books []database.Book) []Book {
	result := make([]Book, len(books))
	for i, book := range books {
		result[i] = DatabaseBookToBook(book)
	}
	return result
}

type Club struct {
	ID        	uuid.UUID 	`json:"id"`
	Name      	string   	`json:"name"`
	Genre  		string		`json:"genre"`
}

func DatabaseClubtoClub(club database.Club) Club {
	return Club{
		ID:        	club.ID,
		Name:      	club.Name,
		Genre:  	club.Genre,	
	}
}

func DatabaseClubsToClubs(clubs []database.Club) []Club {
	result := make([]Club, len(clubs))
	for i, club := range clubs {
		result[i] = DatabaseClubtoClub(club)
	}
	return result
}
