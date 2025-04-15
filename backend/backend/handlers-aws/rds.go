package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"cloudcleanup/utils"
	"github.com/aws/aws-sdk-go-v2/aws"           // ✅ for aws.Bool
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type AWSRequestRDS struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}

type DeleteRequestRDS struct {
	AccessKey    string   `json:"accessKey"`
	SecretKey    string   `json:"secretKey"`
	Region       string   `json:"region"`
	ResourceType string   `json:"resourceType"`
	ResourceIds  []string `json:"resourceIds"` // RDS instance identifiers
}

func GetRDSInstances(w http.ResponseWriter, r *http.Request) {
	var req AWSRequestRDS
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := rds.NewFromConfig(cfg)
	result, err := client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})
	if err != nil {
		http.Error(w, "Failed to fetch RDS instances: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var instanceIDs []string
	for _, db := range result.DBInstances {
		instanceIDs = append(instanceIDs, *db.DBInstanceIdentifier)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"rdsInstances": instanceIDs,
	})
}

func DeleteRDSInstances(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequestRDS
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := rds.NewFromConfig(cfg)
	var deleted []string

	for _, db := range req.ResourceIds {
		_, err := client.DeleteDBInstance(context.TODO(), &rds.DeleteDBInstanceInput{
			DBInstanceIdentifier: &db,
			SkipFinalSnapshot:    aws.Bool(true), // ✅ Proper fix here
		})
		if err == nil {
			deleted = append(deleted, db)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"deletedRDSInstances": deleted,
	})
}
