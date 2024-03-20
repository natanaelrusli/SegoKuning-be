package repository

import (
	"database/sql"
	"errors"

	"github.com/natanaelrusli/segokuning-be/internal/apperror"
	"github.com/natanaelrusli/segokuning-be/internal/model"
)

type UserRepository interface {
	GetUserByPhone(credentials string) (*model.User, error)
	GetUserByEmail(credentials string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	CreateUser(name, email, phone, password string, images_id int) (*model.User, error)
	AddEmail(id int64, email string) error
	AddPhone(id int64, phone string) error
	UpdateUserInfo(id int64, imageId int, name string) error
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrNoUserFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepositoryPostgreSQL) GetUserByEmail(credentials string) (*model.User, error) {
	var user model.User
	query := "SELECT id, email, name, phone, password, name FROM users WHERE email = $1"
	row := ur.db.QueryRow(query, credentials)

	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Phone, &user.Password, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepositoryPostgreSQL) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	query := "SELECT id, email, name, phone, password, name FROM users WHERE id = $1"
	row := ur.db.QueryRow(query, id)

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

func (ur *userRepositoryPostgreSQL) AddEmail(id int64, email string) error {
	query := "UPDATE users SET email = $1 WHERE id = $2"
	_, err := ur.db.Query(query, email, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperror.ErrNoUserFound
		}
		return err
	}

	return nil
}

func (ur *userRepositoryPostgreSQL) AddPhone(id int64, phone string) error {
	query := "UPDATE users SET phone = $1 WHERE id = $2"
	_, err := ur.db.Query(query, phone, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperror.ErrNoUserFound
		}
		return err
	}

	return nil
}

func (ur *userRepositoryPostgreSQL) UpdateUserInfo(id int64, imageId int, name string) error {
	query := "UPDATE users SET images_id = $1, name = $2 WHERE id = $3"
	_, err := ur.db.Exec(query, imageId, name, id)
	if err != nil {
		return err
	}
	return nil
}
