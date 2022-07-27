package middlewares

import (
	"database/sql" // package to encode and decode the json into struct and vice versa
	"fmt"          // models package where User schema is defined
	"log"          // used to access the request and response object of the api
	"os"           // used to read the environment variable

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"        // postgres golang driver
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Form Postgres URL
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic("Connection Refused")
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic("Connection not working properly!")
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
