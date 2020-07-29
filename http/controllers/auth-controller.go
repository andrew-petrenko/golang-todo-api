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
	"io/ioutil"
	"net/http"
	"regexp"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	var errorBag utils.ValidationErrorBag
	validateRegisterUserRequest(&registerRequest, &errorBag)

	if errorBag.ContainsErrors() {
		utils.WriteResponse(w, br.NewResponse(errorBag, false), http.StatusUnprocessableEntity)
		return
	}

	var userRepo repositories.UserRepository
	user := &models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
	if err = userRepo.Store(user); err != nil {
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
	json.Unmarshal(reqBody, &aur)

	var ur repositories.UserRepository
	user, err := ur.FindOneByCriteria("email = ?", aur.Email)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
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

func validateRegisterUserRequest(requestData *RegisterUserRequest, errorBag *utils.ValidationErrorBag) {
	if requestData.Name == "" {
		errorBag.AddError("name", "Name can not be empty")
	}

	if len(requestData.Name) < 2 {
		errorBag.AddError("name", "Name should be at least 2 characters")
	}

	if len(requestData.Name) > 255 {
		errorBag.AddError("name", "Name should be less then 255 characters")
	}

	if requestData.Email == "" {
		errorBag.AddError("email", "Email can not be empty")
	}

	re, _ := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(requestData.Email) || len(requestData.Email) > 255 {
		errorBag.AddError("email", "There is an error in email address")
	}

	if requestData.Password == "" {
		errorBag.AddError("password", "Password can not be empty")
	}

	if len(requestData.Password) < 8 {
		errorBag.AddError("password", "Password must be at least 8 characters")
	}
}
