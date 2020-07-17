package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"type:varchar(255),not_null"`
	Email    string `json:"email" gorm:"type:varchar(100),unique_index,not_null"`
	Password string `json:"-" gorm:"type:varchar(255),not_null"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	u.Password = string(hashedPassword)
}
