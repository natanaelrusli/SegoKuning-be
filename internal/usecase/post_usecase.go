package usecase

import (
	"strings"

	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/repository"
)

type PostUsecase interface {
	CreatePost(userId int64, postInHtml string, tags []string) (*dto.PostData, error)
}

type postUsecaseImpl struct {
	postRepository repository.PostRepository
}

func NewPostUsecaseImpl(postRepository repository.PostRepository) *postUsecaseImpl {
	return &postUsecaseImpl{
		postRepository: postRepository,
	}
}

func (pu *postUsecaseImpl) CreatePost(userId int64, postInHtml string, tags []string) (*dto.PostData, error) {
	// var postData dto.PostData

	stringTags := strings.Join(tags, ", ")

	newPost, err := pu.postRepository.CreateOne(userId, postInHtml, stringTags)
	if err != nil {
		return nil, err
	}

	var newPostData dto.PostData
	newPostData.ID = int64(newPost.ID)
	newPostData.UserId = int64(newPost.UserId)
	newPostData.PostInHtml = newPost.PostInHTML
	newPostData.CreatedAt = newPost.CreatedAt
	newPostData.Tags = tags

	return &newPostData, nil
}
