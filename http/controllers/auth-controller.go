package controllers

import (
	"encoding/json"
	"github.com/andrew-petrenko/golang-todo-api/http/resources"
	br "github.com/andrew-petrenko/golang-todo-api/http/resources/base-response"
	"github.com/andrew-petrenko/golang-todo-api/models"
	"github.com/andrew-petrenko/golang-todo-api/repositories"
	"github.com/andrew-petrenko/golang-todo-api/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
	"io/ioutil"
	"net/http"
)

// TODO add email validation

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"min=1,max=255"`
	Email    string `json:"email" validate:"min=1"`
	Password string `json:"password" validate:"min=8,max=32"`
}

type AuthUserRequest struct {
	Email    string `json:"email" validate:"min=1"`
	Password string `json:"password" validate:"min=8,max=32"`
}

var userRepo repositories.UserRepository

func Register(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Failed to read request body"), http.StatusInternalServerError)
		return
	}

	var registerRequest RegisterUserRequest
	if err := json.Unmarshal(reqBody, &registerRequest); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if err := validator.Validate(registerRequest); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
	if err = userRepo.Create(user); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJwtToken(user.Id)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(resources.AuthenticatedUserResource{
		Id:    user.Id,
		Token: token,
	}, true), http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Failed to read request body"), http.StatusInternalServerError)
		return
	}

	var aur AuthUserRequest
	if err := json.Unmarshal(reqBody, &aur); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	user, err := userRepo.FindOneByCriteria("email = ?", aur.Email)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if err := validator.Validate(aur); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusUnprocessableEntity)
		return
	}

	if user.Id == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("User not found"), http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(aur.Password)); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Incorrect password"), http.StatusForbidden)
		return
	}

	token, err := utils.GenerateJwtToken(user.Id)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(resources.AuthenticatedUserResource{
		Id:    user.Id,
		Token: token,
	}, true), http.StatusOK)
}
