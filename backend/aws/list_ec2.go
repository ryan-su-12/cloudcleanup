package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"cloudcleanup/backend/models" // adjust as per actual module path
)


// Fetch EC2 instances
func FetchEC2Instances(cfg aws.Config, resultChan chan []models.Resource, errorChan chan error) {
	client := ec2.NewFromConfig(cfg)

	// Call DescribeInstances
	resp, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		errorChan <- err
		return
	}

	var resources []models.Resource
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			resources = append(resources, models.Resource{
				ID:   aws.ToString(instance.InstanceId),
				Type: "ec2",
			})
		}
	}

	resultChan <- resources
}

