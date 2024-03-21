package dto

import (
	"mime/multipart"
	"net/url"
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

type LinkEmailRequest struct {
	Email string `json:"email"`
}

type LinkPhoneRequest struct {
	Phone string `json:"Phone"`
}

type ImageUploadRequest struct {
	File *multipart.FileHeader `form:"file"`
	// File       multipart.File        // File represents the uploaded file.
}

type UpdateProfileRequest struct {
	ImageURL string `json:"imageUrl" validate:"required,url"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
}

type FriendQuery struct {
	// validate this
	UserId     int64
	Limit      int    `form:"limit"`
	Offset     int    `form:"offset"`
	OrderBy    string `form:"orderBy"`
	SortBy     string `form:"sortBy"`
	OnlyFriend bool   `form:"onlyFriend"`
	Search     string `form:"search"`
}

func (r *UpdateProfileRequest) Validate() error {
	if r.ImageURL == "" {
		return apperror.ErrImageURLEmpty
	}

	_, err := url.ParseRequestURI(r.ImageURL)
	if err != nil {
		return apperror.ErrInvalidImageURL
	}

	if r.Name == "" {
		return apperror.ErrNoName
	}

	if len(r.Name) < 5 || len(r.Name) > 50 {
		return apperror.ErrInvalidNameLength
	}

	return nil
}

func (r *LinkEmailRequest) Validate() error {
	valid := isValidEmail(r.Email)

	if !valid {
		return apperror.ErrInvalidEmail
	}

	return nil
}

func (r *LinkPhoneRequest) Validate() error {
	valid := isValidPhoneNumber(r.Phone)

	if !valid {
		return apperror.ErrInvalidPhone
	}

	return nil
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

func (r *ImageUploadRequest) Validate() error {
	// Check file type, must in *.jpg | *.jpeg format
	if r.File.Header.Get("Content-Type") != "image/jpeg" && r.File.Header.Get("Content-Type") != "image/jpg" {
		return apperror.ErrImageType
	}

	// Check file size, no more than 2MB, no less than 10KB
	if r.File.Size < 10*1024 {
		return apperror.ErrImageSizeBelow
	}
	if r.File.Size > 2*1024*1024 {
		return apperror.ErrImageSizeAbove
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
