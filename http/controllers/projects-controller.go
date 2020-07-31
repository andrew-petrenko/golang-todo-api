package controllers

import (
	"encoding/json"
	"github.com/andrew-petrenko/golang-todo-api/http/resources"
	br "github.com/andrew-petrenko/golang-todo-api/http/resources/base-response"
	"github.com/andrew-petrenko/golang-todo-api/models"
	"github.com/andrew-petrenko/golang-todo-api/repositories"
	"github.com/andrew-petrenko/golang-todo-api/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"gopkg.in/validator.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type CreateProjectRequest struct {
	Title       string `json:"title" validate:"min=1,max=255"`
	Description string `json:"description" validate:"min=1,max=2000"`
}

type UpdateProjectRequest struct {
	Title       string `json:"title" validate:"min=1,max=255"`
	Description string `json:"description" validate:"min=1,max=2000"`
}

func GetUsersProjects(w http.ResponseWriter, r *http.Request) {
	authToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var claims utils.AuthClaims
	_, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || claims.UserId == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Error"), http.StatusInternalServerError)
		return
	}

	var repo repositories.ProjectRepository
	userProjects, err := repo.FindAllByCriteria("user_id = ?", strconv.Itoa(int(claims.UserId)))
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	var projectResources []resources.ProjectResource
	for _, projectModel := range userProjects {
		projectResources = append(projectResources, resources.ProjectResource{
			ID:          projectModel.Id,
			UserId:      projectModel.UserId,
			Title:       projectModel.Title,
			Description: projectModel.Description,
			CreatedAt:   projectModel.CreatedAt,
			UpdatedAt:   projectModel.UpdatedAt,
		})
	}

	utils.WriteResponse(w, br.NewResponse(projectResources, true), http.StatusOK)
}

func DeleteAllUsersProjects(w http.ResponseWriter, r *http.Request) {
	authToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var claims utils.AuthClaims
	_, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || claims.UserId == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Error"), http.StatusInternalServerError)
		return
	}

	var repo repositories.ProjectRepository
	if err := repo.DeleteAllByCriteria("user_id = ?", strconv.Itoa(int(claims.UserId))); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	authToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var claims utils.AuthClaims
	_, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || claims.UserId == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Error"), http.StatusInternalServerError)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Failed to read request body"), http.StatusInternalServerError)
		return
	}

	var createProjectRequest CreateProjectRequest
	if err := json.Unmarshal(reqBody, &createProjectRequest); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if err := validator.Validate(createProjectRequest); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusUnprocessableEntity)
		return
	}

	project := &models.Project{
		UserId:      claims.UserId,
		Title:       createProjectRequest.Title,
		Description: createProjectRequest.Description,
	}

	var repo repositories.ProjectRepository
	if err := repo.Create(project); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(resources.ProjectResource{
		ID:          project.Id,
		UserId:      project.UserId,
		Title:       project.Title,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}, true), http.StatusCreated)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	authToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var claims utils.AuthClaims
	_, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || claims.UserId == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Error"), http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusBadRequest)
		return
	}

	var repo repositories.ProjectRepository
	project, err := repo.FindOneById(id)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if project.Id == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Project not found"), http.StatusNotFound)
		return
	}

	if project.UserId != claims.UserId {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Forbidden"), http.StatusForbidden)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Failed to read request body"), http.StatusInternalServerError)
		return
	}

	var updateProjectRequest UpdateProjectRequest
	if err := json.Unmarshal(reqBody, &updateProjectRequest); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if err := validator.Validate(updateProjectRequest); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusUnprocessableEntity)
		return
	}

	project.Title = updateProjectRequest.Title
	project.Description = updateProjectRequest.Description

	project, err = repo.Save(project)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(resources.ProjectResource{
		ID:          project.Id,
		UserId:      project.UserId,
		Title:       project.Title,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}, true), http.StatusOK)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	authToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var claims utils.AuthClaims
	_, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || claims.UserId == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Error"), http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusBadRequest)
		return
	}

	var repo repositories.ProjectRepository
	project, err := repo.FindOneById(id)
	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if project.Id == 0 {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Project not found"), http.StatusNotFound)
		return
	}

	if project.UserId != claims.UserId {
		utils.WriteResponse(w, br.NewResponseErrorMessage("Forbidden"), http.StatusForbidden)
		return
	}

	if err := repo.Delete(id); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
