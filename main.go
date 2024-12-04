package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shouryabansal7/BookFam/db"
	"github.com/shouryabansal7/BookFam/handler"
	"github.com/shouryabansal7/BookFam/middleware"
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

	apiCfg, db, err := db.DBSet()
	if err != nil {
		log.Fatal(err)
	}

	// Remember to close the database connection when you're done
	defer db.Close()

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz",handler.HandlerReadiness)
	v1Router.Get("/err",handler.HandlerError)
	v1Router.Post("/users",handler.HandlerCreateUser(apiCfg))
	v1Router.Post("/login", handler.HandlerLogin(apiCfg))
	v1Router.Post("/protected",middleware.MiddlewareAuth(handler.HandlerProtectedRoute,apiCfg))
	v1Router.Post("/book",middleware.MiddlewareAuth(handler.HandlerAddBook,apiCfg))
	v1Router.Get("/books",middleware.MiddlewareAuth(handler.HandlerGetBooks,apiCfg))
	v1Router.Get("/user_books",middleware.MiddlewareAuth(handler.HandlerGetBooksBeingReadByUser,apiCfg))
	v1Router.Post("/clubs",middleware.MiddlewareAuth(handler.HandlerCreateClub,apiCfg))
	v1Router.Post("/club_join/{ClubID}", middleware.MiddlewareAuth(handler.HandlerJoinClub,apiCfg))
	v1Router.Post("/club_leave/{ClubID}", middleware.MiddlewareAuth(handler.HandlerLeaveClub,apiCfg))
	v1Router.Get("/clubs",middleware.MiddlewareAuth(handler.HandlerGetClubs,apiCfg))

	router.Mount("/v1",v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server Starting on port %v", portString)
	err = srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Port:",portString)
}