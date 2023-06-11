package blogController

import (
	errorResponse "blog-platfrom-go/Utils"
	"github.com/gin-gonic/gin"
	"io"
)

type BlogController struct {
}

func (b *BlogController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "200",
	})
}

func (b *BlogController) CreateBlog(c *gin.Context) {
	_, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(418, errorResponse.New(err))
	}
	c.JSON(201, gin.H{
		"response": "201",
	})
}

func New() BlogController {
	return BlogController{}
}
