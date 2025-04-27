package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func FetchRDSInstances(cfg aws.Config, resultChan chan []Resource, errorChan chan error) {
	client := rds.NewFromConfig(cfg)

	resp, err := client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []Resource
	for _, db := range resp.DBInstances {
		resources = append(resources, Resource{
			ID:   aws.ToString(db.DBInstanceIdentifier),
			Type: "rds",
		})
	}

	resultChan <- resources
}
