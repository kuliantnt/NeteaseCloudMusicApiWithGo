package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// 推荐节目
func ProgramRecommend(c *gin.Context) {
	var service service.ProgramRecommendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ProgramRecommend(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
