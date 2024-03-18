package usecase

import (
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/encryptutils"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/jwtutils"
	"github.com/natanaelrusli/segokuning-be/internal/repository"
)

type AuthUsecase interface {
	RegisterUser(name, credentialValue, credentialType, password string) (*dto.UserData, error)
	LoginUser(username, password string) (*dto.LoginRequest, error)
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
	var userData dto.UserData

	hashedPassword, err := au.passwordEncryptor.Hash(password)
	if err != nil {
		return nil, err
	}

	newUser, err := au.userRepository.CreateUser(name, credentialType, credentialValue, hashedPassword)
	if err != nil {
		return nil, err
	}

	if newUser.CredentialType == "email" {
		userData = dto.UserData{
			Email: newUser.CredentialValue,
		}
	} else {
		userData = dto.UserData{
			Phone: newUser.CredentialValue,
		}
	}

	userData.Name = newUser.Name
	userData.AccessToken = "abcabcabc"

	return &userData, nil
}

func (au *authUsecaseImpl) LoginUser(username, password string) (*dto.LoginRequest, error) {
	return nil, nil
}
