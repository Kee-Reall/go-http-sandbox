package blogMiddleware

import (
	blogEntity "blog-platfrom-go/Entities"
	errorResponse "blog-platfrom-go/Utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func BlogValidator(c *gin.Context) {
	var dto blogEntity.BlogInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse.New(err))
		c.Abort()
		return
	} //todo add json Schema validator

	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse.New(err))
		c.Abort()
		return
	}
	c.Next()
	return
}
