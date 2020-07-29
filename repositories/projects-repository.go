package repositories

import (
	"github.com/andrew-petrenko/golang-todo-api/models"
)

type ProjectRepository struct {
	repository
}

func (pr *ProjectRepository) FindAll() ([]models.Project, error) {
	var projects []models.Project
	db, err := pr.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	db.Find(&projects)

	return projects, nil
}

func (pr *ProjectRepository) FindOneById(id int) (*models.Project, error) {
	var project models.Project
	db, err := pr.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.First(&project, id)

	return &project, nil
}

func (pr *ProjectRepository) Delete(id int) error {
	db, err := pr.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Delete(&models.Project{}, id)

	return nil
}
