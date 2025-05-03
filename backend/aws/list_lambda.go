package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"cloudcleanup/backend/models" // adjust as per actual module path
)

func FetchLambdaFunctions(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := lambda.NewFromConfig(cfg)

	resp, err := client.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []models.Resource
	for _, fn := range resp.Functions {
		resources = append(resources, models.Resource{
			ID:   aws.ToString(fn.FunctionName),
			Type: "lambda",
		})
	}

	resultChan <- resources
}


