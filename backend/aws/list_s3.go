package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Fetch S3 buckets
func FetchS3Buckets(cfg aws.Config, resultChan chan []Resource, errorChan chan error) {
	client := s3.NewFromConfig(cfg)

	// Call ListBuckets
	resp, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []Resource
	for _, bucket := range resp.Buckets {
		resources = append(resources, Resource{
			ID:   aws.ToString(bucket.Name),
			Type: "s3",
		})
	}

	resultChan <- resources
}


