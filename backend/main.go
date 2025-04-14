package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("DeleteAWSResources is at: %v\n", DeleteAWSResources)
	r := mux.NewRouter()

	r.HandleFunc("/api/aws/resources", GetAWSResources).Methods("POST")
	r.HandleFunc("/api/aws/delete", DeleteAWSResources).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("404 Not Found: %s %s\n", req.Method, req.URL.Path)
		http.NotFound(w, req)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
