package apperror

import "errors"

// TODO: Group errors and variable name should Err + Entities + Desc
// e.g. Err + Image + InvalidSize = ErrImageInvalidSize

// AUTH
var ErrEmailExists = errors.New("email already exists")
var ErrPhoneExists = errors.New("phone number already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrNoUserFound = errors.New("no user found with the provided credentials")
var ErrInvalidCredentialsType = errors.New("credentialType should be either 'phone' or 'email'")
var ErrInvalidEmail = errors.New("credentialValue should be in email format")
var ErrInvalidPhone = errors.New("credentialValue should be a phone number format")
var ErrInvalidPassword = errors.New("password length should be between 5 and 15 characters")
var ErrInvalidNameLength = errors.New("name length should be between 5 and 50 characters")
var ErrInvalidPasswordLength = errors.New("password length should be between 5 and 50 characters")

var ErrInvalidToken = errors.New("token invalid")

var ErrAlreadyHaveEmail = errors.New("this account already have a linked email address")
var ErrAlreadyHavePhone = errors.New("this account already have a linked phone number")
var ErrImageURLEmpty = errors.New("imageUrl cannot be empty")
var ErrInvalidImageURL = errors.New("imageUrl must be a valid URL")
var ErrNoName = errors.New("name cannot be empty")

// IMAGE
var (
	ErrImageType      = errors.New("file type should be in jgp or jpeg format")
	ErrImageSizeBelow = errors.New("file size should not be less than 10KB")
	ErrImageSizeAbove = errors.New("file size should not be more than 2MB")
)

// POST
var (
	ErrPostInvalidLength = errors.New("post should be between 2 and 500 characters")
	ErrPostTagsEmpty     = errors.New("tags must be filled with at least one item")
)
