package main

import (
	"net/http"
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
	current_day := ""
	utc_time := ""
	github_file_url := ""
	github_repo_url := ""
	status_code := 200

	// Create a JSON response
	response := map[string]interface{}{
		"slack_name":      slack_name,
		"track":           track,
		"current_day":     current_day,
		"utc_time":        utc_time,
		"github_file_url": github_file_url,
		"github_repo_url": github_repo_url,
		"status_code":     status_code,
	}
}
