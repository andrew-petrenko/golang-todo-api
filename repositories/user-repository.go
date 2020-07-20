package repositories

import (
	"github.com/andrew-petrenko/golang-todo-api/models"
)

type UserRepository struct {
	repository
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	db, err := ur.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Create(&user)

	return nil
}

func (ur *UserRepository) GetByCriteria(criteria string, values ...string) ([]models.User, error) {
	db, err := ur.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []models.User
	db.Where(criteria, values).Find(&users)

	return users, err
}
