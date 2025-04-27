package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Resource struct for unified frontend rendering
type Resource struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Fetch EC2 instances
func FetchEC2Instances(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := ec2.NewFromConfig(cfg)

	// Call DescribeInstances
	resp, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []Resource
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			resources = append(resources, Resource{
				ID:   aws.ToString(instance.InstanceId),
				Type: "ec2",
			})
		}
	}

	resultChan <- resources
}

