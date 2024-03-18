package dto

import (
	"errors"
	"regexp"
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

func (u *RegisterRequest) Validate() error {
	if u.CredentialType != "phone" && u.CredentialType != "email" {
		return errors.New("credentialType should be either 'phone' or 'email'")
	}

	if u.CredentialType == "email" {
		if !isValidEmail(u.CredentialValue) {
			return errors.New("credentialValue should be in email format")
		}
	} else {
		if !isValidPhoneNumber(u.CredentialValue) {
			return errors.New("credentialValue should be a phone number format")
		}
	}

	if len(u.Name) < 5 || len(u.Name) > 50 {
		return errors.New("name length should be between 5 and 50 characters")
	}

	if len(u.Password) < 5 || len(u.Password) > 15 {
		return errors.New("password length should be between 5 and 15 characters")
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
