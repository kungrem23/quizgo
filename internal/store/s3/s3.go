package s3

import (
	"context"
	"log"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Connection() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("Loading S3 config error: %v", err)
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	_, err = client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Printf("Loading list buckets error: %v", err)
		return nil, err
	}
	return client, nil
}

func NewS3PresignClient(client *s3.Client) *s3.PresignClient {
	presignClient := s3.NewPresignClient(client)
	return presignClient
}

// func GetImageById(presignClient *s3.PresignClient, id int)
