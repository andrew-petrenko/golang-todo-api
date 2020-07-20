package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"os"
)

type DBConnection struct {
}

func (dbC *DBConnection) GetDB() (*gorm.DB, error) {
	databaseDriver, err := dbC.defineDatabaseDriver()
	if err != nil {
		return nil, err
	}

	connectionString, err := dbC.defineConnectionString(databaseDriver)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(databaseDriver, connectionString)
	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}

	return db, nil
}

func (dbC *DBConnection) defineDatabaseDriver() (string, error) {
	driver := os.Getenv("DB_CONNECTION")
	if driver == "" {
		return "", errors.New("DB_CONNECTION should be defined in .env")
	}

	return driver, nil
}

func (dbC *DBConnection) defineConnectionString(driver string) (string, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.New("DB_HOST should be defined in .env")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.New("DB_PORT should be defined in .env")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return "", errors.New("DB_USER should be defined in .env")
	}

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		return "", errors.New("DB_PASS should be defined in .env")
	}

	database := os.Getenv("DB_DATABASE")
	if database == "" {
		return "", errors.New("DB_DATABASE should be defined in .env")
	}

	switch driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, database), nil
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, database, pass), nil
	case "sqlite3":
		return fmt.Sprintf("/tmp/gorm.db"), nil
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, database), nil
	default:
		return "", errors.New("Invalid database driver")
	}
}
