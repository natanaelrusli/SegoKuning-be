package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
}

func NewFriendHandler() *FriendHandler {
	return &FriendHandler{}
}

func (fh *FriendHandler) GetFriendList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get friend list",
	})
}

func (fh *FriendHandler) AddFriend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Add friend",
	})
}

func (fh *FriendHandler) RemoveFriend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete friend",
	})
}
