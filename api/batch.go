package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

func Batch(c *gin.Context) {
	var service service.BatchService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Batch(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
