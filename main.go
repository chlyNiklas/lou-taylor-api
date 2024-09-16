package main

import (
	"log"
	"net/http"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/config"
	"github.com/chlyNiklas/lou-taylor-api/database"
	"github.com/chlyNiklas/lou-taylor-api/middleware"
	"github.com/chlyNiklas/lou-taylor-api/server_service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := config.New()
	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	service := server_service.New(cfg, db)
	mw := middleware.New(cfg)

	r := mux.NewRouter()

	// Attach the generated handler using your server implementation
	apiHandler := api.Handler(api.NewStrictHandler(service, nil))

	r.PathPrefix("/").Handler(apiHandler)
	r.Use(mw.Authentication)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
