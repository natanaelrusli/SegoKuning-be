package handler

import (
	"net/http"
	"strconv"

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
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)

	if limit == 0 {
		limit = 5
	}

	res, err := fh.friendUsecase.GetFriendList(1, int64(limit), int64(offset))
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
			Limit:  int(limit),
			Offset: int(offset),
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
