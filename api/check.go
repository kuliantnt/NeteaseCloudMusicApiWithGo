package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

func CheckMusic(c *gin.Context) {
	var service service.CheckMusicService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CheckMusic(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
