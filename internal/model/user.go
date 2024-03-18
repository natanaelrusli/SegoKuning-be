package model

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	CredentialType  string `json:"credentialType"`
	CredentialValue string `json:"credentialValue"`
	Password        string `json:"password"`
}
