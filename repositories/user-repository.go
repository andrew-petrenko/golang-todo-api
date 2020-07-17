package repositories

import (
	"github.com/andrew-petrenko/golang-todo-api/models"
	"github.com/andrew-petrenko/golang-todo-api/utils"
)

type UserRepository struct {
	connector *utils.DbConnection
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	db, err := ur.connector.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Create(&user)

	return nil
}

func (ur *UserRepository) GetByCriteria(criteria string, values ...string) ([]models.User, error) {
	db, err := ur.connector.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []models.User
	db.Where(criteria, values).Find(&users)

	return users, err
}
