package dto

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
