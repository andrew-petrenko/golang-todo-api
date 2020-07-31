package main

import (
	"github.com/andrew-petrenko/golang-todo-api/http"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

// TODO think about some Core package with Authenticated User information
// TODO think about responses, error handling and status codes

func main() {
	if err := http.InitHttpServer(); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
