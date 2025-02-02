package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GithubURL       string `json:"github_url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application")
	res := Response{
		Email:           "atanda0x@gmail.com",
		CurrentDateTime: time.Now().UTC().Format(time.RFC3339),
		GithubURL:       "https/github.com/atanda0x/hng_task_0",
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Encoding res: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := "8000"
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
