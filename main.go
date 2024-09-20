package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/chlyNiklas/lou-taylor-api/api"
	"github.com/chlyNiklas/lou-taylor-api/authentication"

	"github.com/chlyNiklas/lou-taylor-api/config"
	"github.com/chlyNiklas/lou-taylor-api/controller"
	"github.com/chlyNiklas/lou-taylor-api/database"
	"github.com/chlyNiklas/lou-taylor-api/image_service"
	"github.com/gorilla/mux"
)

const configFile = "./config.toml"

func main() {

	cfg := config.Default()
	cfg.ReadFlags()
	if err := cfg.ReadFile(cfg.ConfigPath); err != nil {
		log.Println(err)
	}

	fmt.Println(cfg.TOML())
	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	img := image_service.New(cfg.Images)

	spec, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}
	auth := authentication.New(cfg.Authentication, spec)

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
