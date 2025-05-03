package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"cloudcleanup/backend/models" // adjust as per actual module path
)

func FetchRDSInstances(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := rds.NewFromConfig(cfg)

	resp, err := client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []models.Resource
	for _, db := range resp.DBInstances {
		resources = append(resources, models.Resource{
			ID:   aws.ToString(db.DBInstanceIdentifier),
			Type: "rds",
		})
	}

	resultChan <- resources
}
