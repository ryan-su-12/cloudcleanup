package handlersaws

import (

	"encoding/json"
	"net/http"

	"cloudcleanup/backend/utils"  // LoadAWSConfig function
	"cloudcleanup/backend/models"  
	"cloudcleanup/backend/aws"    // FetchEC2Instances, FetchS3Buckets, etc.
)


type AWSRequestCommon struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}


// ListAllAWSResources launches goroutines to fetch resources
func ListAllAWSResources(w http.ResponseWriter, r *http.Request) {
	var req AWSRequestCommon
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)

	if err != nil {
		http.Error(w, "AWS config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resultChan := make(chan []models.Resource)
	errorChan := make(chan error)

	// Launch concurrent fetchers
	go aws.FetchEC2Instances(cfg, resultChan, errorChan)
	go aws.FetchS3Buckets(cfg, resultChan, errorChan)
	go aws.FetchLambdaFunctions(cfg, resultChan, errorChan)
	go aws.FetchRDSInstances(cfg, resultChan, errorChan)
	go aws.FetchDynamoDBTables(cfg, resultChan, errorChan)
	go aws.FetchSQSQueues(cfg, resultChan, errorChan)
	go aws.FetchSNSTopics(cfg, resultChan, errorChan)

	totalFetchers := 7
	var allResources []models.Resource

	for i := 0; i < totalFetchers; i++ {
		select {
		case res := <-resultChan:
			allResources = append(allResources, res...)
		case err := <-errorChan:
			println("Fetch error:", err.Error())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResources)
}
