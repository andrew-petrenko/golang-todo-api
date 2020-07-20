package http

import (
	"fmt"
	"github.com/andrew-petrenko/golang-todo-api/http/router"
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
