package repository

import (
	"database/sql"

	"github.com/natanaelrusli/segokuning-be/internal/model"
)

type FriendRepository interface {
	GetFriendsByUserID(userId int64, limit int64, offset int64) ([]model.FriendUserData, error)
}

type friendRepositoryImpl struct {
	db *sql.DB
}

func NewFriendRepository(db *sql.DB) *friendRepositoryImpl {
	return &friendRepositoryImpl{
		db: db,
	}
}

func (fr *friendRepositoryImpl) GetFriendsByUserID(userId int64, limit int64, offset int64) ([]model.FriendUserData, error) {
	var users []model.FriendUserData

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
		group by 
			users.id,
			users.name,
			i.url,
			users.images_id,
			users.created_at
		LIMIT $1
		OFFSET $2;
		`

	rows, err := fr.db.Query(query, limit, offset)
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
