package main

import (
	"os"

	"github.com/te-fa-bene/api-go/internal/app/http/routes"
)

// @title 			Te fa Bene API
// @version 		0.1
// @description MVP for a restaurant API
// @BasePath 		/api/v1
func main() {
	r := routes.NewRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
