package models

type Project struct {
	BaseModel
	UserId      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignkey:UserRefer"`
	Title       string `json:"title" gorm:"varchar(100);not_null"`
	Description string `json:"description" gorm:"varchar(255);not_null"`
}
