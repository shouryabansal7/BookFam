package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shouryabansal7/BookFam/db"
	"github.com/shouryabansal7/BookFam/token"
)

type login_cred struct{
	Email string
	Password string
}

func HandlerLogin(apiCfg *db.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		fmt.Printf("The Request body is %v\n", r.Body)
		var u login_cred
		json.NewDecoder(r.Body).Decode(&u)
		//fmt.Printf("The user request value %v", u)

		db_u, err := apiCfg.DB.GetUserByEmail(r.Context(),u.Email)
		//fmt.Println(db_u.Email);

		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No username found")
		}

		if db_u.Password == u.Password {
			tokenString, err := token.CreateToken(u.Email, db_u.ID)
			if err!=nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Errorf("Error while generating token")
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, tokenString)
			return
		}else{
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid credentials")
		}
	}
}
