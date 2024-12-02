package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/shouryabansal7/BookFam/handler"
)

func main(){
	fmt.Println("Hello World")

	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz",handler.HandlerReadiness)
	v1Router.Get("/err",handler.HandlerError)
	router.Mount("/v1",v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server Starting on port %v", portString)
	err := srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Port:",portString)
}