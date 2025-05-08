package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketBasics struct {
	S3Client *s3.Client
}

func NewBucketBasics() (*BucketBasics, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-southeast-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"",
			"",
			"",
		),
		))
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return &BucketBasics{S3Client: client}, nil
}

func (basics *BucketBasics) UploadFile(bucketName string, objectKey string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
		return err
	}
	defer file.Close()
	_, err = basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
			fileName, bucketName, objectKey, err)
	}
	return err
}
