package models

import "time"

type BaseModel struct {
	Id        uint       `json:"id" gorm:"column:id; primary_key"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
