package apperror

import "errors"

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

var ErrImageType = errors.New("file type should be in jgp or jpeg format")
var ErrImageSizeBelow = errors.New("file size should not be less than 10KB")
var ErrImageSizeAbove = errors.New("file size should not be more than 2MB")
