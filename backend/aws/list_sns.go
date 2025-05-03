package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"cloudcleanup/backend/models" // adjust as per actual module path
)

func FetchSNSTopics(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := sns.NewFromConfig(cfg)

	resp, err := client.ListTopics(context.TODO(), &sns.ListTopicsInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []models.Resource
	for _, topic := range resp.Topics {
		resources = append(resources, models.Resource{
			ID:   aws.ToString(topic.TopicArn),
			Type: "sns",
		})
	}

	resultChan <- resources
}
