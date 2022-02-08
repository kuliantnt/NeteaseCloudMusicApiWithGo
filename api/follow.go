package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

func Follow(c *gin.Context) {
	var service service.FollowService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Follow(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
