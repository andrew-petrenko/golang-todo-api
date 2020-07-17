package main

import (
	"github.com/andrew-petrenko/golang-todo-api/http"
	"log"
)

func main() {
	if err := http.InitHttpServer(); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
