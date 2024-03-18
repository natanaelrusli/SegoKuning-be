package model

type Profile struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Name        string `json:"name"`
	ImageURL    string `json:"imageUrl"`
	FriendsID   []int  `json:"friendsId"`
	FriendCount int    `json:"friendCount"`
	CreatedAt   string `json:"createdAt"`
}
