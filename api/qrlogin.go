package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// LoginQrkey 二维码生成接口
//  @param c
func LoginQrkey(c *gin.Context) {
	var service service.LoginQrKeyService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginQrkey(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
