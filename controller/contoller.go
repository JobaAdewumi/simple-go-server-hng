package controller

import (
	"fmt"
	"net/http"
	"server/model"

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

	user_id := context.Param("userId")

	if user_id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

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
	user_id := context.Param("userId")

	if user_id == "" {
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
		// ID:   user_id,
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
