package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secretKey = []byte("secret-key")

type Claims struct {
	UserID uuid.UUID `json:"id"`
	Email  string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(username string, id uuid.UUID) (string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id":id,
		"username":username,
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString,err := token.SignedString(secretKey)
	if err!=nil{
		return "",err
	}

	return tokenString,nil
}

func VerifyToken(tokenString string)error{
	token,err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{}, error){
		return secretKey,nil
	})

	if err!=nil{
		return err
	}
	if !token.Valid{
		return fmt.Errorf("invalid token")
	}
	return nil
}

func ExtractJWTInfo(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC (you can change this depending on your JWT signature algorithm)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// If there's an error in parsing, return it
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Extract the claims from the token if it's valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid claims")

}