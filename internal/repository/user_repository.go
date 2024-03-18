package repository

import (
	"database/sql"

	"github.com/natanaelrusli/segokuning-be/internal/model"
)

type UserRepository interface {
	GetUserByPhone(credentials string) (*model.User, error)
	CreateUser(name, email, phone, password string, images_id int) (*model.User, error)
}

type userRepositoryPostgreSQL struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepositoryPostgreSQL {
	return &userRepositoryPostgreSQL{
		db: db,
	}
}

func (ur *userRepositoryPostgreSQL) GetUserByPhone(credentials string) (*model.User, error) {
	var user model.User
	query := "SELECT id, email, name, phone, password, name FROM users WHERE phone = $1"
	row := ur.db.QueryRow(query, credentials)

	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Phone, &user.Password, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepositoryPostgreSQL) CreateUser(name, email, phone, password string, images_id int) (*model.User, error) {
	var newUser model.User
	query := "INSERT INTO users (email, phone, password, name, images_id) VALUES ($1, $2, $3, $4, $5) RETURNING phone, email, name"
	err := ur.db.QueryRow(query, email, phone, password, name, images_id).Scan(
		&newUser.Phone,
		&newUser.Email,
		&newUser.Name,
	)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
