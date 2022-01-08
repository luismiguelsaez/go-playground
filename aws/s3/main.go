package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatalf("Error loading default config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	response, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Error while getting bucket list: %v", err)
	}

	for _, bucket := range response.Buckets {
		log.Printf("Bucket name: %s", *bucket.Name)
	}
}
