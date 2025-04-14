package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	// ✅ Apply CORS middleware globally
	r.Use(enableCORS)
	log.Println("✅ CORS Middleware is now active")

	// 🧪 Simple health check route
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("✅ /ping endpoint hit")
		w.Write([]byte("pong"))
	}).Methods("GET")

	// 🔗 Core AWS routes
	r.HandleFunc("/api/aws/resources", GetAWSResources).Methods("POST")
	r.HandleFunc("/api/aws/delete", DeleteAWSResources).Methods("POST")

	// ✅ Catch-all OPTIONS handler for preflight requests
	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("✅ Global OPTIONS handler triggered for:", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	})

	// 🚫 Fallback for unknown paths
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("404 Not Found: %s %s\n", req.Method, req.URL.Path)
		http.NotFound(w, req)
	})

	// 🚀 Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// 🌐 CORS Middleware
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
