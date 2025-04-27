package main

import (
	"log"
	"net/http"


	"cloudcleanup/backend/handlersaws"
	"github.com/gorilla/mux"


)

func main() {

	
	r := mux.NewRouter().StrictSlash(true)

	
	r.Use(enableCORS)
	log.Println("✅ CORS Middleware is now active")

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("✅ /ping endpoint hit")
		w.Write([]byte("pong"))
	}).Methods("GET")
	/*
	// centralizing the routes that each aws service can have 
	r.HandleFunc("/api/aws/resources", handlers.GetEC2Instances).Methods("POST")
	r.HandleFunc("/api/aws/delete", handlers.DeleteEC2Instances).Methods("POST")


	r.HandleFunc("/api/aws/s3/resources", handlers.GetS3Buckets).Methods("POST")
	r.HandleFunc("/api/aws/s3/delete", handlers.DeleteS3Buckets).Methods("POST")

	r.HandleFunc("/api/aws/rds/resources", handlers.GetRDSInstances).Methods("POST")
	r.HandleFunc("/api/aws/rds/delete", handlers.DeleteRDSInstances).Methods("POST")

	r.HandleFunc("/api/aws/lambda/resources", handlers.GetLambdaFunctions).Methods("POST")
	r.HandleFunc("/api/aws/lambda/delete", handlers.DeleteLambdaFunctions).Methods("POST")

	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("✅ Global OPTIONS handler triggered for:", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	})
		*/

	r.HandleFunc("/api/aws/list-resources", handlers.ListAWSResources).Methods("POST")


	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("404 Not Found: %s %s\n", req.Method, req.URL.Path)
		http.NotFound(w, req)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}


// Getting perms from CORS for GO
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("🛡️ CORS middleware hit for:", r.Method, r.URL.Path)

		// Allow requests from Next.js frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
