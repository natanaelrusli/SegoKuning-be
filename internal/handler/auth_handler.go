package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/segokuning-be/internal/apperror"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ah.authUsecase.RegisterUser(req.Name, req.CredentialValue, req.CredentialType, req.Password)

	if err != nil && err.Error() == "email already exists" || err != nil && err.Error() == "phone number already exists" {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := ah.authUsecase.LoginUser(req.CredentialValue, req.CredentialType, req.Password)

	if err != nil && err.Error() == "no user found with the provided credentials" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err != nil && err.Error() == "invalid credentials" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

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

func (ah *AuthHandler) LinkEmail(c *gin.Context) {
	var req dto.LinkEmailRequest
	userId := c.Value("ctx-user-id").(int64)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ah.authUsecase.LinkEmail(userId, req.Email)
	if err != nil && err == apperror.ErrAlreadyHaveEmail {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil && err == apperror.ErrEmailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "link email success",
	})
}

func (ah *AuthHandler) LinkPhone(c *gin.Context) {
	var req dto.LinkPhoneRequest
	userId := c.Value("ctx-user-id").(int64)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ah.authUsecase.LinkPhone(userId, req.Phone)
	if err != nil && err == apperror.ErrAlreadyHaveEmail {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil && err == apperror.ErrEmailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "link phone success",
	})
}
