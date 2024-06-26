package repository

import (
	"database/sql"
	"fmt"

	"github.com/natanaelrusli/segokuning-be/internal/apperror"
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/model"
)

type FriendRepository interface {
	GetFriendsByUserID(friendQuery dto.FriendQuery) ([]model.FriendUserData, error)
	AddFriendByUserID(userId, targetUserId int64) error
	DeleteFriendByUserID(userId, targetUserId int64) error
}

type friendRepositoryImpl struct {
	db *sql.DB
}

func NewFriendRepository(db *sql.DB) *friendRepositoryImpl {
	return &friendRepositoryImpl{
		db: db,
	}
}

func (fr *friendRepositoryImpl) GetFriendsByUserID(friendQuery dto.FriendQuery) ([]model.FriendUserData, error) {
	var users []model.FriendUserData
	var args []interface{}

	query := `
		select 
			users.id,
			users.name,
			i.url as image_url,
			COUNT(uid2) AS friend_count,
			users.created_at
		from 
			friendships as f
		inner join users on users.id  = f.uid1 or users.id = f.uid2  
		inner join images i on users.images_id = i.id
		where
			users.id is not null
		`

	if friendQuery.OnlyFriend {
		args = append(args, int64(friendQuery.UserId))
		query += fmt.Sprintf(" AND (f.uid1 = $%d OR f.uid2 = $%d)", len(args), len(args))
	}

	if friendQuery.Search != "" {
		args = append(args, friendQuery.Search)
		query += fmt.Sprintf(" AND users.name ILIKE '%%' || $%d || '%%'", len(args))
	}

	query += `
		group by 
			users.id,
			users.name,
			i.url,
			users.images_id,
			users.created_at
	`

	if friendQuery.SortBy == "friendCount" {
		query += " ORDER BY friend_count"
	} else if friendQuery.SortBy == "createdAt" {
		query += " ORDER BY users.created_at"
	} else {
		query += " ORDER BY users.created_at" // Default sorting by createdAt
	}

	if friendQuery.OrderBy == "asc" {
		query += " ASC"
	} else {
		query += " DESC" // Default ordering is DESC
	}

	args = append(args, int64(friendQuery.Limit))
	query += fmt.Sprintf(" LIMIT $%d", len(args))

	args = append(args, int64(friendQuery.Offset))
	query += fmt.Sprintf(" OFFSET $%d", len(args))

	rows, err := fr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.FriendUserData
		if err := rows.Scan(&user.ID, &user.Name, &user.ImageURL, &user.FriendCount, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (fr *friendRepositoryImpl) AddFriendByUserID(userId, targetUserId int64) error {
	var count int

	queryCheck := "SELECT COUNT(*) FROM friendships WHERE (uid1 = $1 AND uid2 = $2) OR (uid1 = $2 AND uid2 = $1)"
	err := fr.db.QueryRow(queryCheck, userId, targetUserId).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return apperror.ErrAlreadyFriend
	}

	queryInsert := "INSERT INTO friendships (uid1, uid2) VALUES ($1, $2)"
	_, err = fr.db.Exec(queryInsert, userId, targetUserId)
	if err != nil {
		return err
	}

	return nil
}

func (fr *friendRepositoryImpl) DeleteFriendByUserID(userId, targetUserId int64) error {
	var count int
	queryCheck := "SELECT COUNT(*) FROM friendships WHERE (uid1 = $1 AND uid2 = $2) OR (uid1 = $2 AND uid2 = $1)"
	err := fr.db.QueryRow(queryCheck, userId, targetUserId).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return apperror.ErrNotFriend
	}

	query := "DELETE from friendships WHERE (uid1 = $1 and uid2 = $2) or (uid1 = $2 and uid2 = $1);"
	_, err = fr.db.Exec(query, targetUserId, userId)

	return err
}
