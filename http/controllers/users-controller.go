package controllers

import (
	"encoding/json"
	"github.com/andrew-petrenko/golang-todo-api/http/resources"
	br "github.com/andrew-petrenko/golang-todo-api/http/resources/base-response"
	"github.com/andrew-petrenko/golang-todo-api/repositories"
	"github.com/andrew-petrenko/golang-todo-api/utils"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type UpdateUserInfoRequest struct {
	Name string `json:"name" validate:"min=1,max=255"`
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	authToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var claims utils.AuthClaims
	_, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || claims.UserId == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Error"), http.StatusInternalServerError)
		return
	}

	var repo repositories.UserRepository

	user, err := repo.FindOneById(claims.UserId)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(resources.UserPersonalInfoResource{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, true), http.StatusOK)
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Failed to read request body"), http.StatusInternalServerError)
		return
	}

	var updateUserRequest UpdateUserInfoRequest
	if err := json.Unmarshal(reqBody, &updateUserRequest); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	authToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var claims utils.AuthClaims
	_, err = jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || claims.UserId == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Error"), http.StatusInternalServerError)
		return
	}

	var repo repositories.UserRepository
	user, err := repo.FindOneById(claims.UserId)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	user.Name = updateUserRequest.Name

	user, err = repo.Save(user)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(user, true), http.StatusOK)
}
