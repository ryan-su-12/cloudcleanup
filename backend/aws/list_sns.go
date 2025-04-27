package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func FetchSNSTopics(cfg aws.Config, resultChan chan []Resource, errorChan chan error) {
	client := sns.NewFromConfig(cfg)

	resp, err := client.ListTopics(context.TODO(), &sns.ListTopicsInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []Resource
	for _, topic := range resp.Topics {
		resources = append(resources, Resource{
			ID:   aws.ToString(topic.TopicArn),
			Type: "sns",
		})
	}

	resultChan <- resources
}
