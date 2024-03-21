package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

// NOTE-CLEAN4: Handler is on adapter layer, is used for handling http request, extract data (e.g. json, form body), doing necessary validation, and then delegate the actual business logic to the use case layer
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

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// err := req.Validate()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	log.Print("before upload image")
	res, err := ih.imageUsecase.UploadImage(req.File)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error uploading image, please try again",
		})
	}

	c.JSON(http.StatusCreated, res)
}
