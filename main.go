package main

import (
	blogController "blog-platfrom-go/Controllers"
	blogMiddleware "blog-platfrom-go/Middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	blogCtrl := blogController.New()
	blogs := r.Group("/blogs")
	{
		blogs.GET("/")
		blogs.GET("/:id", blogCtrl.GetOne)
		blogs.POST("/", blogMiddleware.BlogValidator, blogCtrl.CreateBlog)
		blogs.PUT("/:id")
		blogs.DELETE("/:id")
	}

	r.GET("/", func(context *gin.Context) {
		context.JSON(201, gin.H{
			"response": "201",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
