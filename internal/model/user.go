package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	ImagesID  int       `json:"images_id"`
	CreatedAt time.Time `json:"created_at"`
}
