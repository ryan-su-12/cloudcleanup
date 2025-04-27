package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func FetchLambdaFunctions(cfg aws.Config, resultChan chan []Resource, errorChan chan error) {
	client := lambda.NewFromConfig(cfg)

	resp, err := client.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []Resource
	for _, fn := range resp.Functions {
		resources = append(resources, Resource{
			ID:   aws.ToString(fn.FunctionName),
			Type: "lambda",
		})
	}

	resultChan <- resources
}
