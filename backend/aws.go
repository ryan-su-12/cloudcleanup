package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Struct to receive credentials from frontend
type AWSRequest struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}

type DeleteRequest struct {
	AccessKey   string   `json:"accessKey"`
	SecretKey   string   `json:"secretKey"`
	Region      string   `json:"region"`
	InstanceIDs []string `json:"instanceIds"`
}

func GetAWSResources(w http.ResponseWriter, r *http.Request) {
	// Decode incoming request
	var req AWSRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("JSON Decode error:", err) // 👈 Add this line for debugging
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Load AWS config with static credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(req.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(req.AccessKey, req.SecretKey, ""),
		),
	)
	if err != nil {
		http.Error(w, "Failed to load AWS config: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Call DescribeInstances
	output, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		http.Error(w, "Failed to describe instances: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Collect instance IDs
	var instanceIDs []string
	for _, reservation := range output.Reservations {
		for _, instance := range reservation.Instances {
			instanceIDs = append(instanceIDs, *instance.InstanceId)
		}
	}

	// Respond with JSON
	json.NewEncoder(w).Encode(map[string]interface{}{
		"instances": instanceIDs,
	})
}

func DeleteAWSResources(w http.ResponseWriter, r *http.Request) {

	// Decode the request
	var req DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(req.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(req.AccessKey, req.SecretKey, ""),
		),
	)
	if err != nil {
		http.Error(w, "Failed to load AWS config: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Prepare terminate input
	input := &ec2.TerminateInstancesInput{
		InstanceIds: req.InstanceIDs,
	}

	// Call TerminateInstances
	output, err := ec2Client.TerminateInstances(context.TODO(), input)
	if err != nil {
		http.Error(w, "Failed to terminate instances: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Build response
	var terminated []string
	for _, inst := range output.TerminatingInstances {
		terminated = append(terminated, *inst.InstanceId)
	}

	// Return result
	json.NewEncoder(w).Encode(map[string]interface{}{
		"terminated": terminated,
	})
	
}
