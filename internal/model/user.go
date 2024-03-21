package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Password  string       `json:"password"`
	ImagesID  int          `json:"images_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type FriendUserData struct {
	ID          int       `json:"userId"`
	Name        string    `json:"name"`
	ImageURL    *string   `json:"imageUrl"`
	FriendCount int64     `json:"friendCount"`
	CreatedAt   time.Time `json:"createdAt"`
}
