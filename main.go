package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Working Fine!"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", healthCheckHandler).Methods("GET")

	fmt.Println("Server started Listening on :8000")
	fmt.Println("health-check api :", "http://localhost:8000/health-check")

	// CORS middleware
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
