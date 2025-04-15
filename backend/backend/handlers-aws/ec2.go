package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"cloudcleanup/utils" // Update to match your module name
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type AWSRequestEC2 struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}

type DeleteRequestEC2 struct {
	AccessKey    string   `json:"accessKey"`
	SecretKey    string   `json:"secretKey"`
	Region       string   `json:"region"`
	ResourceType string   `json:"resourceType"`
	ResourceIds  []string `json:"resourceIds"` // Instance IDs
}


func GetEC2Instances(w http.ResponseWriter, r *http.Request) {
	var req AWSRequestEC2
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := ec2.NewFromConfig(cfg)
	result, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		http.Error(w, "Failed to fetch EC2 instances: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var instanceIDs []string
	for _, res := range result.Reservations {
		for _, inst := range res.Instances {
			instanceIDs = append(instanceIDs, *inst.InstanceId)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"instances": instanceIDs,
	})
}

func DeleteEC2Instances(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequestEC2
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := ec2.NewFromConfig(cfg)
	input := &ec2.TerminateInstancesInput{
		InstanceIds: req.ResourceIds,
	}

	output, err := client.TerminateInstances(context.TODO(), input)
	if err != nil {
		http.Error(w, "Terminate failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var terminated []string
	for _, inst := range output.TerminatingInstances {
		terminated = append(terminated, *inst.InstanceId)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"terminated": terminated,
	})
}

