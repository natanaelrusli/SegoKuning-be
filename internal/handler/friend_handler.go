package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/usecase"
)

type FriendHandler struct {
	friendUsecase usecase.FriendUsecase
}

func NewFriendHandler(friendUsecase usecase.FriendUsecase) *FriendHandler {
	return &FriendHandler{
		friendUsecase: friendUsecase,
	}
}

func (fh *FriendHandler) GetFriendList(c *gin.Context) {
	var friendQuery dto.FriendQuery

	err := c.ShouldBindQuery(&friendQuery)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	friendQuery.UserId = c.Value("ctx-user-id").(int64)

	if friendQuery.Limit == 0 {
		friendQuery.Limit = 5
	}

	if friendQuery.Offset == 0 {
		friendQuery.Offset = 0
	}

	if friendQuery.SortBy == "" {
		friendQuery.SortBy = "createdAt"
	}

	if friendQuery.OrderBy == "" {
		friendQuery.OrderBy = "desc"
	}

	res, err := fh.friendUsecase.GetFriendList(friendQuery)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	paginationRes := dto.PaginationResponse{
		Message: res.Messagge,
		Data:    res.Data,
		Meta: dto.PaginationMeta{
			Limit:  int(friendQuery.Limit),
			Offset: int(friendQuery.Offset),
			Total:  len(res.Data),
		},
	}

	c.JSON(http.StatusOK, paginationRes)
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
