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