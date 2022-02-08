package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

func ResourceLike(c *gin.Context) {
	var service service.ResourceLikeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ResourceLike(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
