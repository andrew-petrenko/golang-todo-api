package resources

import "time"

type ProjectResource struct {
	ID          uint       `json:"id"`
	UserId      uint       `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
