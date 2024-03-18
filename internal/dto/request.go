package dto

import (
	"regexp"

	"github.com/natanaelrusli/segokuning-be/internal/apperror"
)

type RegisterRequest struct {
	CredentialType  string `json:"credentialType"`
	CredentialValue string `json:"credentialValue"`
	Name            string `json:"name"`
	Password        string `json:"password"`
}

type LoginRequest struct {
	CredentialType  string `json:"credentialType"`
	CredentialValue string `json:"credentialValue"`
	Password        string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	if r.CredentialType != "phone" && r.CredentialType != "email" {
		return apperror.ErrInvalidCredentials
	}

	if r.CredentialType == "email" {
		if !isValidEmail(r.CredentialValue) {
			return apperror.ErrInvalidEmail
		}
	} else {
		if !isValidPhoneNumber(r.CredentialValue) {
			return apperror.ErrInvalidPhone
		}
	}

	if len(r.Password) < 5 || len(r.Password) > 15 {
		return apperror.ErrInvalidPassword
	}

	return nil
}

func (u *RegisterRequest) Validate() error {
	if u.CredentialType != "phone" && u.CredentialType != "email" {
		return apperror.ErrInvalidCredentialsType
	}

	if u.CredentialType == "email" {
		if !isValidEmail(u.CredentialValue) {
			return apperror.ErrInvalidEmail
		}
	} else {
		if !isValidPhoneNumber(u.CredentialValue) {
			return apperror.ErrInvalidPhone
		}
	}

	if len(u.Name) < 5 || len(u.Name) > 50 {
		return apperror.ErrInvalidNameLength
	}

	if len(u.Password) < 5 || len(u.Password) > 15 {
		return apperror.ErrInvalidPasswordLength
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

func isValidPhoneNumber(phoneNumber string) bool {
	phoneNumberRegex := `^\+\d{7,13}$`
	return regexp.MustCompile(phoneNumberRegex).MatchString(phoneNumber)
}
