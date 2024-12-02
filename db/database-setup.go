package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/shouryabansal7/BookFam/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func DBSet() (*ApiConfig, *sql.DB, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := &ApiConfig{
		DB: dbQueries,
	}

	return apiCfg, db, nil
}

