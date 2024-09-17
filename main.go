package main

import (
	"log"
	"net/http"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/authentication"
	"github.com/chlyNiklas/lou-taylor-api/config"
	"github.com/chlyNiklas/lou-taylor-api/controller"
	"github.com/chlyNiklas/lou-taylor-api/database"
	"github.com/chlyNiklas/lou-taylor-api/image_service"
	"github.com/gorilla/mux"
)

func main() {

	cfg := config.New()
	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	img := image_service.New(cfg.Images)

	spec, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}
	auth := authentication.New(cfg.Security, spec)

	controller := controller.New(cfg, db, img, auth)

	r := mux.NewRouter()

	// Attach the generated handler using your server implementation
	apiHandler := api.Handler(api.NewStrictHandler(controller, nil))

	r.PathPrefix("/").Handler(apiHandler)
	r.Use(auth.Authentication)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
