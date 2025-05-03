package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"cloudcleanup/backend/models" // adjust as per actual module path
)

func FetchDynamoDBTables(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := dynamodb.NewFromConfig(cfg)

	resp, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []models.Resource
	for _, tableName := range resp.TableNames {
		resources = append(resources, models.Resource{
			ID:   tableName,
			Type: "dynamodb",
		})
	}

	resultChan <- resources
}
