package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api", handlerFunc)
	listenPort := ":80"
	http.ListenAndServe(listenPort, nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Parse GET parameters
	url := r.URL.Query()
	slack_name := url.Get("slack_name")
	track := url.Get("track")

	// Configure response variables
	utc_time := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	current_day := time.Now().Weekday().String()
	github_file_url := "https://github.com/obiMadu/hngX-stage1/blob/2dde12e9795fa1a3858086b169aea0d804c17018/main.go"
	github_repo_url := "https://github.com/obiMadu/hngX-stage1"
	status_code := 200

	// Create a JSON response
	response := map[string]interface{}{
		"slack_name":      slack_name,
		"current_day":     current_day,
		"utc_time":        utc_time,
		"track":           track,
		"github_file_url": github_file_url,
		"github_repo_url": github_repo_url,
		"status_code":     status_code,
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(jsonResponse)
}
