package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// 心动模式播放列表
func PlaymodeIntelligenceList(c *gin.Context) {
	var service service.PlaymodeIntelligenceListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaymodeIntelligenceList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
