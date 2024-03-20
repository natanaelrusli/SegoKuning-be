package usecase

import (
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/repository"
)

type FriendUsecase interface {
	GetFriendList(userId, limit, offset int64) (*dto.FriendsResponse, error)
}

type friendUsecaseImpl struct {
	friendRepository repository.FriendRepository
	userRepository   repository.UserRepository
}

func NewFriendUsecaseImpl(
	friendRepository repository.FriendRepository,
	userRepository repository.UserRepository,
) *friendUsecaseImpl {
	return &friendUsecaseImpl{
		friendRepository: friendRepository,
		userRepository:   userRepository,
	}
}

func (fu *friendUsecaseImpl) GetFriendList(userId, limit, offset int64) (*dto.FriendsResponse, error) {
	var res dto.FriendsResponse

	users, err := fu.friendRepository.GetFriendsByUserID(userId, limit, offset)
	if err != nil {
		return nil, err
	}

	res.Messagge = "success"
	res.Data = users

	return &res, nil
}
