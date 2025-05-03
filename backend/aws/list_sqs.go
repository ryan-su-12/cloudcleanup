package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"cloudcleanup/backend/models" // adjust as per actual module path
)

func FetchSQSQueues(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := sqs.NewFromConfig(cfg)

	resp, err := client.ListQueues(context.TODO(), &sqs.ListQueuesInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []models.Resource
	if resp.QueueUrls != nil {
		for _, url := range resp.QueueUrls {
			resources = append(resources, models.Resource{
				ID:   url,
				Type: "sqs",
			})
		}
	}

	resultChan <- resources
}
