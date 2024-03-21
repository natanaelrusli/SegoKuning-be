package model

import "time"

type Post struct {
	ID         int       `json:"postId"`
	UserId     int       `json:"userId"`
	PostInHTML string    `json:"postInHTML"`
	Tags       string    `json:"tags"`
	CreatedAt  time.Time `json:"createdAt"` // Assuming ISO 8601 format as string
}
