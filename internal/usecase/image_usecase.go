package usecase

import (
	"mime/multipart"

	"github.com/natanaelrusli/segokuning-be/internal/dto"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/aws"
	"github.com/natanaelrusli/segokuning-be/internal/repository"
)

// NOTE-CLEAN3: Usecase layer contains the business logic of the application and orchestrates interactions between different parts of the system, such as handling user requests, perform validations, and coordinating data flow between the repository and other components

// business logic is rules, requirements and constraints specific to our application domain or business requirements, it represents the core functionality and behaviour that drives our application and distinguishes it from others

// the logic is difference from the base logic or validation in the handler, e.g. request key, request content (string, number), file type, etc

// 1 same as repository
type ImageUsecase interface {
	UploadImage(file *multipart.FileHeader) (*dto.ImageData, error)
}

// 2 struct for implementation of the usecase
// concrete implementation struct
type imageUsecaseImpl struct {
	imageRepository repository.ImageRepository
}

// 3 constructor function (factory) creating an instance of the imageUsecaseImpl
func NewImageUsecaseImpl(
	imageRepository repository.ImageRepository,
) *imageUsecaseImpl {
	return &imageUsecaseImpl{
		imageRepository: imageRepository,
	}
}

// 4 this is the main usecase for UploadImage
// here it accept params from handler and pass to repository to be processed in the data store (e.g. db)
// then return the result (domain model) and use DTO
func (iu *imageUsecaseImpl) UploadImage(file *multipart.FileHeader) (*dto.ImageData, error) {

	urlStr, err := aws.ImageUpload(file)
	if err != nil {
		return nil, err
	}

	// TODO: apanihh
	// var imageData dto.ImageData
	// because in this the data send to repository and get back from repository is the same, I think its okay to skip?

	newImage, err := iu.imageRepository.CreateOne(urlStr)
	if err != nil {
		return nil, err
	}

	var newImageData dto.ImageData
	newImageData.URL = newImage.URL

	return &newImageData, nil
}
