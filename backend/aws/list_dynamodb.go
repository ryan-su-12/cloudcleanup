package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func FetchDynamoDBTables(cfg aws.Config, resultChan chan []Resource, errorChan chan error) {
	client := dynamodb.NewFromConfig(cfg)

	resp, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []Resource
	for _, tableName := range resp.TableNames {
		resources = append(resources, Resource{
			ID:   tableName,
			Type: "dynamodb",
		})
	}

	resultChan <- resources
}
