package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/usecase"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := req.Validate()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(req)

	res, err := ah.authUsecase.RegisterUser(req.Name, req.CredentialValue, req.CredentialType, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	user, token, err := ah.authUsecase.LoginUser(req.CredentialValue, req.CredentialType, req.Password)

	if err != nil {
		c.JSON(500, err.Error())

		return
	}

	loginUserData := dto.LoginUserData{
		Email:       user.Email,
		Name:        user.Name,
		Phone:       user.Phone,
		AccessToken: token,
	}

	data := dto.LoginResponse{
		Message: "User logged successfully",
		Data:    loginUserData,
	}
	c.JSON(200, data)
}
