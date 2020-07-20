package controllers

import (
	br "github.com/andrew-petrenko/golang-todo-api/http/resources/base-response"
	"github.com/andrew-petrenko/golang-todo-api/repositories"
	"github.com/andrew-petrenko/golang-todo-api/utils"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var repo repositories.ProjectRepository
	projects, err := repo.GetAll()

	if err != nil {
		utils.WriteResponse(w, br.NewResponse(err.Error(), false), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(projects, true), http.StatusOK)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, br.NewResponse(nil, true), http.StatusNotImplemented)
}

func GetOneProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		utils.WriteResponse(w, br.NewResponse(err.Error(), false), http.StatusBadRequest)
		return
	}

	var repo repositories.ProjectRepository
	project, err := repo.GetById(id)
	if err != nil {
		utils.WriteResponse(w, br.NewResponse(err.Error(), false), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, br.NewResponse(project, true), http.StatusOK)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, br.NewResponse(nil, false), http.StatusNotImplemented)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		utils.WriteResponse(w, br.NewResponse(err.Error(), false), http.StatusBadRequest)
		return
	}

	var repo repositories.ProjectRepository
	if err := repo.Delete(id); err != nil {
		utils.WriteResponse(w, br.NewResponse(err.Error(), false), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
