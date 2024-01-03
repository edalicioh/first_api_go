package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edalicioh/first_api_go/configs"
	"github.com/edalicioh/first_api_go/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	err = configs.InitDB()
	if err != nil {
		log.Fatalf("Error init db: %s", err)
	}

	router := mux.NewRouter()

	routes.RegisterRoutesApi(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is listening on port %s\n", port)

	err = http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
