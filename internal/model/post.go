package model

type Post struct {
	ID         int      `json:"id"`
	PostInHTML string   `json:"postInHTML"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"createdAt"` // Assuming ISO 8601 format as string
}
