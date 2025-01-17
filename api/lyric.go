package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// 获取歌词
func Lyric(c *gin.Context) {
	var service service.LyricService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Lyric(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
