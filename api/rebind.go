package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// 重新绑定手机
func Rebind(c *gin.Context) {
	var service service.RebindService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Rebind(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
