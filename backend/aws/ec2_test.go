package aws_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"cloudcleanup/backend/models"
	cleanaws "cloudcleanup/backend/aws"
)

// Mock client struct
type mockEC2Client struct {
	DescribeInstancesOutput *ec2.DescribeInstancesOutput
	Error                   error
}

func (m *mockEC2Client) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.DescribeInstancesOutput, nil
}
