package repositories

import (
	"github.com/andrew-petrenko/golang-todo-api/models"
)

type UserRepository struct {
	repository
}

func (ur *UserRepository) Create(user *models.User) error {
	db, err := ur.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Create(&user)

	return nil
}

func (ur *UserRepository) Save(user *models.User) (*models.User, error) {
	db, err := ur.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.Save(&user)

	return user, nil
}

func (ur *UserRepository) FindOneById(id uint) (*models.User, error) {
	db, err := ur.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	db.First(&user, id)

	return &user, nil
}

func (ur *UserRepository) FindOneByCriteria(criteria string, values ...string) (*models.User, error) {
	db, err := ur.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	db.Where(criteria, values).First(&user)

	return &user, nil
}
