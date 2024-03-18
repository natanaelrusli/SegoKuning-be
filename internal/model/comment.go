package model

type Comment struct {
	ID        int    `json:"id"`
	PostsID   int    `json:"postsId"`
	UserID    int    `json:"userId"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"createdAt"`
}
