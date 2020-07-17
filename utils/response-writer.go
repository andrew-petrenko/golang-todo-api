package utils

import (
	"encoding/json"
	br "github.com/andrew-petrenko/golang-todo-api/http/resources/base-response"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, response *br.Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
