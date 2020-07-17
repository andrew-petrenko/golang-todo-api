package repositories

import (
	"github.com/andrew-petrenko/golang-todo-api/http/resources"
	"github.com/andrew-petrenko/golang-todo-api/utils"
)

//TODO why the hell I return resources?

type ProjectRepository struct {
	connector *utils.DbConnection
}

func (pr *ProjectRepository) GetAll() (*[]resources.Project, error) {
	var projects []resources.Project
	db, err := pr.connector.GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	db.Find(&projects)

	return &projects, nil
}

func (pr *ProjectRepository) GetById(id int) (*[]resources.Project, error) {
	var project []resources.Project
	db, err := pr.connector.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.First(&project, id)

	return &project, nil
}

func (pr *ProjectRepository) Delete(id int) error {
	db, err := pr.connector.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Delete(&resources.Project{}, id)

	return nil
}
