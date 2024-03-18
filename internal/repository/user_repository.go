package repository

import (
	"database/sql"

	"github.com/natanaelrusli/segokuning-be/internal/model"
)

type UserRepository interface {
	GetUser(credentials string) error
	CreateUser(name, credentialType, credentialValue, password string) (*model.User, error)
}

type userRepositoryPostgreSQL struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepositoryPostgreSQL {
	return &userRepositoryPostgreSQL{
		db: db,
	}
}

func (ur *userRepositoryPostgreSQL) GetUser(credentials string) error {
	query := "SELECT * FROM users WHERE credentialValue = $1"
	row := ur.db.QueryRow(query, credentials)

	var id int
	var name, credentialType, credentialValue, password string

	err := row.Scan(&id, &name, &credentialType, &credentialValue, &password)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepositoryPostgreSQL) CreateUser(name, credentialType, credentialValue, password string) (*model.User, error) {
	var newUser model.User
	query := "INSERT INTO users (name, credentialType, credentialValue, password) VALUES ($1, $2, $3, $4) RETURNING credentialValue, name"
	err := ur.db.QueryRow(query, name, credentialType, credentialValue, password).Scan(
		&newUser.CredentialValue,
		&newUser.Name,
	)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
