package main

import (
	"log"

	"github.com/DimaSU2020/car_rental_service/config"
	"github.com/DimaSU2020/car_rental_service/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("dotenv not loaded: %v", err)
	}

	c := config.MustNewCfg("./config.json")

	srv, cleanup, err := app.BuildHTTPServer(c)
	if err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
	defer cleanup()

	log.Fatal(srv.ListenAndServe())
}
