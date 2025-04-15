package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"cloudcleanup/utils"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type AWSRequestLambda struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}

type DeleteRequestLambda struct {
	AccessKey    string   `json:"accessKey"`
	SecretKey    string   `json:"secretKey"`
	Region       string   `json:"region"`
	ResourceType string   `json:"resourceType"`
	ResourceIds  []string `json:"resourceIds"` // Lambda function names
}

func GetLambdaFunctions(w http.ResponseWriter, r *http.Request) {
	var req AWSRequestLambda
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := lambda.NewFromConfig(cfg)
	result, err := client.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{})
	if err != nil {
		http.Error(w, "Failed to fetch Lambda functions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var functionNames []string
	for _, fn := range result.Functions {
		functionNames = append(functionNames, *fn.FunctionName)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"lambdaFunctions": functionNames,
	})
}

func DeleteLambdaFunctions(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequestLambda
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cfg, err := utils.LoadAWSConfig(req.AccessKey, req.SecretKey, req.Region)
	if err != nil {
		http.Error(w, "Config error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := lambda.NewFromConfig(cfg)
	var deleted []string

	for _, fn := range req.ResourceIds {
		_, err := client.DeleteFunction(context.TODO(), &lambda.DeleteFunctionInput{
			FunctionName: &fn,
		})
		if err == nil {
			deleted = append(deleted, fn)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"deletedLambdaFunctions": deleted,
	})
}
