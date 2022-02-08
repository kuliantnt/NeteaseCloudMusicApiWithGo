package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// 转发动态
func EventForward(c *gin.Context) {
	var service service.EventForwardService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EventForward(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 删除动态
func EventDel(c *gin.Context) {
	var service service.EventDelService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EventDel(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取动态列表（网页版）
func Event(c *gin.Context) {
	var service service.EventService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Event(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
