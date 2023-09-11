package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/api", returnInfo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type simpleMessage struct {
	SlackName     string `json:"slack_name"`
	UtcTime       string `json:"utc_time"`
	CurrentDay    string `json:"current_day"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	Track         string `json:"track"`
	StatusCode    int    `json:"status_code"`
}

func returnInfo(c *gin.Context) {

	// Get query parameters from url
	slack_name := c.Request.URL.Query().Get("slack_name")
	if slack_name == "" {
		slack_name = "Joba Adewumi"
	}
	track := c.Request.URL.Query().Get("track")
	if track == "" {
		track = "Backend"
	}

	// get current day and utc date
	currentUtcTime := time.Now().UTC()
	currentDay := time.Now().Weekday()

	simpleMessage := simpleMessage{
		SlackName:     slack_name,
		UtcTime:       currentUtcTime.Format(time.RFC3339),
		CurrentDay:    currentDay.String(),
		GithubFileUrl: "https://github.com/JobaAdewumi/simple-go-server-hng/blob/main/server.go",
		GithubRepoUrl: "https://github.com/JobaAdewumi/simple-go-server-hng",
		Track:         track,
		StatusCode:    200,
	}

	c.IndentedJSON(http.StatusOK, simpleMessage)
}
