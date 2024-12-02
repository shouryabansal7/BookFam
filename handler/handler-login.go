package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shouryabansal7/BookFam/token"
)

type login_cred struct{
	Email string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	fmt.Printf("The Request body is %v\n", r.Body)
	var u login_cred
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)

    if u.Email == "chek" && u.Password == "123456"{
		tokenString, err := token.CreateToken(u.Email)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No username found")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	}else{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}
