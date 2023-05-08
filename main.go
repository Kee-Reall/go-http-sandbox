package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type BlogInput struct {
	Name        string `json:"name" validate:"required,max=15,min=1"`
	Description string `json:"description" validate:"required,min=1,max=500"`
	WebsiteUrl  string `json:"websiteUrl" validate:"required,url"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(201, gin.H{
			"response": "201",
		})
	})
	r.POST("/", func(context *gin.Context) {
		var json BlogInput
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"response": "fuck you",
			})
			return
		}

		validate := validator.New()
		if err := validate.Struct(json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"response": err.Error(),
			})
			return
		}

		fmt.Printf("%v+", &json)
		context.JSON(http.StatusOK, gin.H{"response": "ok"})
		return
	})

	err := r.Run()
	if err != nil {
		return
	}
}
