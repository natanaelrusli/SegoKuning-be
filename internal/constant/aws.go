package constant

import (
	"os"

	"github.com/natanaelrusli/segokuning-be/internal/config"
)

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
}

func InitAwsConfig() AWSConfig {
	config.GetEnv()

	region := "ap-southeast-1"
	accessKeyID := os.Getenv("S3_ID")
	secretAccessKey := os.Getenv("S3_SECRET_KEY")
	bucket := os.Getenv("S3_BUCKET_NAME")

	return AWSConfig{
		Region:          region,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Bucket:          bucket,
	}

}
