package usecase

import (
	"github.com/natanaelrusli/segokuning-be/internal/apperror"
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/model"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/encryptutils"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/jwtutils"
	"github.com/natanaelrusli/segokuning-be/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	RegisterUser(name, credentialValue, credentialType, password string) (*dto.UserData, error)
	LoginUser(credentials, credentialType, password string) (*model.User, string, error)
}

type authUsecaseImpl struct {
	userRepository    repository.UserRepository
	passwordEncryptor encryptutils.PasswordEncryptor
	jwtUtil           jwtutils.JWTUtil
}

func NewAuthUsecaseImpl(
	userRepository repository.UserRepository,
	passwordEncryptor encryptutils.PasswordEncryptor,
	jwtUtil jwtutils.JWTUtil,
) *authUsecaseImpl {
	return &authUsecaseImpl{
		userRepository:    userRepository,
		passwordEncryptor: passwordEncryptor,
		jwtUtil:           jwtUtil,
	}
}

func (au *authUsecaseImpl) RegisterUser(name, credentialValue, credentialType, password string) (*dto.UserData, error) {
	// check user exists
	var res *model.User
	var err error

	if credentialType == "email" {
		res, err = au.userRepository.GetUserByEmail(credentialValue)
		if res != nil && err == nil {
			return nil, apperror.ErrEmailExists
		}
	} else {
		res, err = au.userRepository.GetUserByPhone(credentialValue)
		if res != nil && err == nil {
			return nil, apperror.ErrPhoneExists
		}
	}

	var userData dto.UserData

	hashedPassword, err := au.passwordEncryptor.Hash(password)
	if err != nil {
		return nil, err
	}

	if credentialType == "email" {
		userData = dto.UserData{
			Email: credentialValue,
		}
	} else {
		userData = dto.UserData{
			Phone: credentialValue,
		}
	}

	newUser, err := au.userRepository.CreateUser(name, userData.Email, userData.Phone, hashedPassword, 0)
	if err != nil {
		return nil, err
	}

	if credentialType == "email" {
		userData = dto.UserData{
			Email: newUser.Email,
		}
	} else {
		userData = dto.UserData{
			Phone: newUser.Phone,
		}
	}

	userData.Name = newUser.Name
	userData.AccessToken = "abcabcabc"

	return &userData, nil
}

func (au *authUsecaseImpl) LoginUser(credentials, credentialType, password string) (*model.User, string, error) {
	var res *model.User
	var err error

	if credentialType == "email" {
		res, err = au.userRepository.GetUserByEmail(credentials)
	} else {
		res, err = au.userRepository.GetUserByPhone(credentials)
	}

	if err != nil {
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil {
		return nil, "", apperror.ErrInvalidCredentials
	}

	token, err := au.jwtUtil.Sign(int64(res.ID))
	if err != nil {
		return nil, "", err
	}

	return res, token, nil
}
