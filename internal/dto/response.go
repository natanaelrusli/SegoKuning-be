package dto

import "time"

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
	// ID  int64  `json:"id"`
	URL string `json:"imageUrl"`
}

type PostData struct {
	ID         int64     `json:"id"`
	UserId     int64     `json:"userId"`
	PostInHtml string    `json:"postInHtml"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreatePostResponse struct {
	Message string   `json:"message"`
	Data    PostData `json:"data"`
}

type UploadImageResponse struct {
	Message string    `json:"message"`
	Data    ImageData `json:"data"`
}
