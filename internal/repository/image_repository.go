package repository

import (
	"database/sql"

	"github.com/natanaelrusli/segokuning-be/internal/model"
)

// NOTE-CLEAN2: Repository layer responsible for abstracting the details of how data is persisted and retrieved from the data store (e.g. database)
// basically first we need to create a struct to connect to the database (or data store) provider (number 2)
// then we need to create a function that will do some execution in the database, e.g. create user, get product, etc (below)
// then we store all the available execution function (or method) in the SomethingRepository interface (see above / number 1)

// 1: this export what the repository can do (actionable items)
// e.g. CreateUser, GetUserById, etc
// TODO: still don't know the connection between the interface and methods below
type ImageRepository interface {
	CreateOne(url string) (*model.Image, error) // this get error when not match with what in usecase
}

// 2: struct for postgre DB
type imageRepositoryPostgreSQL struct {
	db *sql.DB
}

// 3: instantiate repository to connect the db
func NewImageRepository(db *sql.DB) *imageRepositoryPostgreSQL {
	return &imageRepositoryPostgreSQL{
		db: db,
	}
}

// we declare this function using method defining receiver in Go. the method receiver is `ir *imageRepositoryPostgreSQL` which indicates that `CreateOne` method is belong to the imageRepositoryPostgreSQL
// note if you forget, we use pointer type *imageRepository... because we want to get the memory address and change the actual variable, because in Go, when we pass variable it makes a copy of the params
func (ir *imageRepositoryPostgreSQL) CreateOne(url string) (*model.Image, error) {
	var newImage model.Image
	query := "INSERT INTO images (url) VALUES ($1) RETURNING id"
	// QueryRow execute query and expect to return at most one row
	// Scan copies the columns from the matched row into the variable
	err := ir.db.QueryRow(query, url).Scan(&newImage.ID, &newImage.URL)

	if err != nil {
		return nil, err
	}

	return &newImage, nil
}
