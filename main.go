package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"src/sampleCRUD/controllers"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := controllers.EmployeeController()
	fmt.Printf("Starting server on the port %s...\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
