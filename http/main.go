package http

import (
	"fmt"
	"github.com/andrew-petrenko/golang-todo-api/http/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func InitHttpServer() error {
	connectionString := parseEnvFile()
	if err := http.ListenAndServe(connectionString, router.InitRouter()); err != nil {
		return err
	}

	return nil
}

func parseEnvFile() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("APP_HOST")
	if host == "" {
		log.Fatal("APP_HOST can not be empty")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("APP_PORT can not be empty")
	}

	return fmt.Sprintf("%s:%s", host, port)
}
