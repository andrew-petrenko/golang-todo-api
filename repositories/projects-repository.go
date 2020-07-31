package repositories

import (
	"github.com/andrew-petrenko/golang-todo-api/models"
)

type ProjectRepository struct {
	repository
}

func (pr *ProjectRepository) Create(project *models.Project) error {
	db, err := pr.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Create(&project)

	return nil
}

func (pr *ProjectRepository) Save(project *models.Project) (*models.Project, error) {
	db, err := pr.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.Save(&project)

	return project, nil
}

func (pr *ProjectRepository) FindAllByCriteria(criteria string, values ...string) ([]models.Project, error) {
	var projects []models.Project
	db, err := pr.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.Where(criteria, values).Find(&projects)

	return projects, nil
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

func (pr *ProjectRepository) DeleteAllByCriteria(criteria string, values ...string) error {
	db, err := pr.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Delete(models.Project{}, criteria, values)

	return nil
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
