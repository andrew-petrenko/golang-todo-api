package http

import (
	"fmt"
	"github.com/andrew-petrenko/golang-todo-api/http/router"
	"log"
	"net/http"
	"os"
)

func InitHttpServer() error {
	connectionString := defineConnectionString()
	if err := http.ListenAndServe(connectionString, router.InitRouter()); err != nil {
		return err
	}

	return nil
}

func defineConnectionString() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("APP_PORT can not be empty")
	}

	return fmt.Sprintf(":%s", port)
}
