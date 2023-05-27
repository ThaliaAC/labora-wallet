package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresDBHandler struct {
	*sql.DB
}

func LoadEnvVar() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	return host, port, dbName, user, password
}

func Connect_DB() (*PostgresDBHandler, error) {

	host, port, dbName, user, password := LoadEnvVar()
	psqlInfo := fmt.Sprintf("host=%s port=%s dbName=%s user=%s password=%s sslmode=disable", host, port, dbName, user, password)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	fmt.Println("Succesful connection to database:", dbConn)
	DbHandler := &PostgresDBHandler{dbConn}
	var result int
	err = dbConn.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The request to the database is active")
	return DbHandler, err
}
