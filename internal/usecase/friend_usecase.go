package usecase

import (
	"github.com/natanaelrusli/segokuning-be/internal/apperror"
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/repository"
)

type FriendUsecase interface {
	GetFriendList(friendQuery dto.FriendQuery) (*dto.FriendsResponse, error)
	AddFriend(userId, targetId int64) error
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

func (fu *friendUsecaseImpl) GetFriendList(friendQuery dto.FriendQuery) (*dto.FriendsResponse, error) {
	var res dto.FriendsResponse

	users, err := fu.friendRepository.GetFriendsByUserID(friendQuery)
	if err != nil {
		return nil, err
	}

	res.Messagge = "success"
	res.Data = users

	return &res, nil
}

func (fu *friendUsecaseImpl) AddFriend(userId, targetId int64) error {
	user, err := fu.userRepository.GetUserByID(targetId)
	if userId == targetId {
		return apperror.ErrAddSelf
	}

	if user == nil {
		return apperror.ErrNoUserFound
	}

	if err != nil {
		return err
	}

	err = fu.friendRepository.AddFriendByUserID(userId, targetId)
	if err != nil {
		return err
	}

	return nil
}
