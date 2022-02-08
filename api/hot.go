package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// 热门话题
func HotTopic(c *gin.Context) {
	var service service.HotTopicService
	if err := c.ShouldBind(&service); err == nil {
		res := service.HotTopic(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
