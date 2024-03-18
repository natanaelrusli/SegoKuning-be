package aws

// functions name:
// 1. newSessionAWS: create a new session for AWS
// 2. uploadFile: upload file to s3 bucket

import (
	"log"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/natanaelrusli/segokuning-be/internal/constant"
)

func initAWSSession() (*session.Session, error) {
	awsConfig := constant.InitAwsConfig()
	region := awsConfig.Region
	accessKeyID := awsConfig.AccessKeyID
	secretAccessKey := awsConfig.SecretAccessKey

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"",
		),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

func uploadS3(uploader *s3manager.Uploader, file multipart.File, bucketName string, fileName string) (string, error) {
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		ACL:    aws.String("public-read"),
		Body:   file,
	})

	return res.Location, err
}

func ImageUpload(sourceFile *multipart.FileHeader) (string, error) {
	sess, err := initAWSSession()
	if err != nil {
		log.Println("Failed to create AWS session:", err)
		return "", err
	}

	file, err := sourceFile.Open()
	if err != nil {
		log.Println("Failed to open file:", err)
		return "", err
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + sourceFile.Filename
	bucketName := constant.InitAwsConfig().Bucket
	urlStr, err := uploadS3(uploader, file, bucketName, fileName)

	if err != nil {
		return "", err
	}
	return urlStr, nil
}
