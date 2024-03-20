package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/usecase"
)

// struct for image handler
type ImageHandler struct {
	imageUsecase usecase.ImageUsecase
}

// same like repository and usecase
func NewImageHandler(imageUsecase usecase.ImageUsecase) *ImageHandler {
	return &ImageHandler{
		imageUsecase: imageUsecase,
	}
}

func (ih *ImageHandler) UploadImage(c *gin.Context) {
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	log.Printf("Failed to get file: %v", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Please provide a file to upload, and the type should be in jpg or jpeg format",
	// 	})
	// 	return
	// }

	var req dto.ImageUploadRequest

	if err := c.BindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ih.imageUsecase.UploadImage(req.File)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error uploading image, please try again",
		})
	}

	c.JSON(http.StatusCreated, res)
}
