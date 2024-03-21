package repository

import (
	"database/sql"

	"github.com/natanaelrusli/segokuning-be/internal/model"
)

type PostRepository interface {
	CreateOne(userId int64, postInHtml, tags string) (*model.Post, error)
}

type postRepositoryPostgreSQL struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *postRepositoryPostgreSQL {
	return &postRepositoryPostgreSQL{
		db: db,
	}
}

func (pr *postRepositoryPostgreSQL) CreateOne(userId int64, postInHtml string, tags string) (*model.Post, error) {
	var newPost model.Post
	query := "INSERT INTO posts (user_id, post_content, tags) VALUES ($1, $2, $3) RETURNING id, user_id, post_content, tags, created_at"
	// tags convert to string
	err := pr.db.QueryRow(query, userId, postInHtml, tags).Scan(&newPost.ID, &newPost.UserId, &newPost.PostInHTML, &newPost.Tags, &newPost.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &newPost, nil
}
