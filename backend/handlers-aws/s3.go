package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"cloudcleanup/utils"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSRequest struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}

type DeleteRequest struct {
	AccessKey    string   `json:"accessKey"`
	SecretKey    string   `json:"secretKey"`
	Region       string   `json:"region"`
	ResourceType string   `json:"resourceType"`
	ResourceIds  []string `json:"resourceIds"` // Bucket names
}

func GetS3Buckets(w http.ResponseWriter, r *http.Request) {
	var req AWSRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := s3.NewFromConfig(cfg)
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		http.Error(w, "Failed to fetch S3 buckets: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var bucketNames []string
	for _, bucket := range result.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"buckets": bucketNames,
	})
}

func DeleteS3Buckets(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := s3.NewFromConfig(cfg)
	var deleted []string

	for _, bucket := range req.ResourceIds {
		_, err := client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
			Bucket: &bucket,
		})
		if err == nil {
			deleted = append(deleted, bucket)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"deletedBuckets": deleted,
	})
}
