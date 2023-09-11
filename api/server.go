package main

// Import all used packages
import (
	"fmt"
	"net/http"
	"time"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	// Get query parameters from url
	slack_name := r.URL.Query().Get("slack_name")
	if slack_name == "" {
		slack_name = "Joba Adewumi"
	}
	track := r.URL.Query().Get("track")
	if track == "" {
		track = "Backend"
	}

	// get current day and utc date
	currentUtcTime := time.Now().UTC()
	currentDay := time.Now().Weekday()

	// Response back in json format
	fmt.Fprintf(w, `{"slack_name": "%[3]s", "current_day": "%[1]s", "utc_time": "%[2]s", "track": "%[4]s", "github_file_url": "https", "github_repo_url": "https", "status_code": 200}`, currentDay, currentUtcTime, slack_name, track)
}

// Testing
// func Main() {
// 	http.Handle("/api", http.HandlerFunc(handleFunc))
// }

// Main function
// func Main() {
// 	// Api route
// 	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
// 		// Get query parameters from url
// 		slack_name := r.URL.Query().Get("slack_name")
// 		if slack_name == "" {
// 			slack_name = "Joba Adewumi"
// 		}
// 		track := r.URL.Query().Get("track")
// 		if track == "" {
// 			track = "Backend"
// 		}

// 		// get current day and utc date
// 		currentUtcTime := time.Now().UTC()
// 		currentDay := time.Now().Weekday()

// 		// Response back in json format
// 		fmt.Fprintf(w, `{"slack_name": "%[3]s", "current_day": "%[1]s", "utc_time": "%[2]s", "track": "%[4]s", "github_file_url": "https", "github_repo_url": "https", "status_code": 200}`, currentDay, currentUtcTime, slack_name, track)
// 	})

// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
