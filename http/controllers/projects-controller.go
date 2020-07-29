package controllers

import (
	"github.com/andrew-petrenko/golang-todo-api/http/resources"
	br "github.com/andrew-petrenko/golang-todo-api/http/resources/base-response"
	"github.com/andrew-petrenko/golang-todo-api/repositories"
	"github.com/andrew-petrenko/golang-todo-api/utils"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var repo repositories.ProjectRepository
	projects, err := repo.FindAll()

	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	var projectResources []resources.ProjectResource
	for _, project := range projects {
		projectResources = append(projectResources, resources.ProjectResource{
			ID:          project.Id,
			UserId:      project.UserId,
			Title:       project.Title,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		})
	}

	utils.WriteResponse(w, br.NewResponse(projectResources, true), http.StatusOK)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, br.NewResponse(nil, true), http.StatusNotImplemented)
}

func GetOneProject(w http.ResponseWriter, r *http.Request) {
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

	utils.WriteResponse(w, br.NewResponse(resources.ProjectResource{
		ID:          project.Id,
		UserId:      project.UserId,
		Title:       project.Title,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}, true), http.StatusOK)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, br.NewResponse(nil, false), http.StatusNotImplemented)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusBadRequest)
		return
	}

	var repo repositories.ProjectRepository
	if err := repo.Delete(id); err != nil {
		utils.WriteResponse(w, br.NewResponseErrorMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
