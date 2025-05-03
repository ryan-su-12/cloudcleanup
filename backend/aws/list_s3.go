package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"cloudcleanup/backend/models" // adjust as per actual module path
)

// Fetch S3 buckets
func FetchS3Buckets(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := s3.NewFromConfig(cfg)

	// Call ListBuckets
	resp, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []models.Resource
	for _, bucket := range resp.Buckets {
		resources = append(resources, models.Resource{
			ID:   aws.ToString(bucket.Name),
			Type: "s3",
		})
	}

	resultChan <- resources
}


