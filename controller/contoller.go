package controller

import (
	"fmt"
	"net/http"
	"server/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(context *gin.Context) {
	var input model.NameInput

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Empty string"})
		return
	}
	fmt.Print(uuid.New().ID())
	person := model.Person{
		ID:   int(uuid.New().ID()),
		Name: input.Name,
	}

	savedPerson, err := person.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if savedPerson.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Rows Created": savedPerson})
}

func Read(context *gin.Context) {

	user_id, err := strconv.Atoi(context.Param("userId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	// if user_id == "" {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
	// 	return
	// }

	person, err := model.FindPersonById(user_id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if person.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"person": person})

}

func Update(context *gin.Context) {
	user_id, err := strconv.Atoi(context.Param("userId"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}
	var input model.NameInput

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errFunc := model.UpdatePersonById(user_id)

	if errFunc != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errFunc.Error()})
		return
	}

	newPerson := model.Person{
		ID:   user_id,
		Name: input.Name,
	}

	savedPerson, err := newPerson.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Rows Affected": savedPerson})
}

func Delete(context *gin.Context) {

	user_id := context.Param("userId")

	if user_id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	string, err := model.DeletePersonById(user_id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if string == "Error Deleting" {

		context.JSON(http.StatusBadRequest, gin.H{"error": string})
		return
	}
	if string == "Person cannot be found" {

		context.JSON(http.StatusBadRequest, gin.H{"error": string})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": string})

}

// func main() {
// 	router := gin.Default()
// 	router.Use(CORSMiddleware())
// 	router.GET("/api", returnInfo)

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	if err := router.Run(":" + port); err != nil {
// 		log.Panicf("error: %s", err)
// 	}
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, UPDATE, DELETE")

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
