package dto

import "github.com/natanaelrusli/segokuning-be/internal/model"

type PaginationMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type PaginationResponse struct {
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

type UserRegistrationResponse struct {
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}

type UserData struct {
	Email       string `json:"email,omitempty"`
	Name        string `json:"name"`
	Phone       string `json:"phone,omitempty"`
	AccessToken string `json:"accessToken"`
}

type LoginUserData struct {
	Email       string `json:"email,omitempty"`
	Name        string `json:"name"`
	Phone       string `json:"phone,omitempty"`
	AccessToken string `json:"accessToken"`
}

type LoginResponse struct {
	Message string        `json:"message"`
	Data    LoginUserData `json:"data"`
}

type ImageData struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type FriendsResponse struct {
	Messagge string                 `json:"message"`
	Data     []model.FriendUserData `json:"data"`
}
