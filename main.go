package main

import (
	"github.com/andrew-petrenko/golang-todo-api/http"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	if err := http.InitHttpServer(); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
