package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"os"
)

type DbConnection struct {
}

func (dbC *DbConnection) GetConnection() (*gorm.DB, error) {
	connectionString, err := dbC.parseEnvFile()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}

	return db, nil
}

func (dbC *DbConnection) parseEnvFile() (string, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return "", errors.New("DB_HOST should be defined in .env")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return "", errors.New("DB_PORT should be defined in .env")
	}

	dbUsername := os.Getenv("DB_USER")
	if dbUsername == "" {
		return "", errors.New("DB_USER should be defined in .env")
	}

	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		return "", errors.New("DB_PASS should be defined in .env")
	}

	dbDatabase := os.Getenv("DB_DATABASE")
	if dbDatabase == "" {
		return "", errors.New("DB_DATABASE should be defined in .env")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPass, dbHost, dbPort, dbDatabase), nil
}
