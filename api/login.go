package api

import (
	"github.com/lianlian/singo/service"

	"github.com/gin-gonic/gin"
)

// LoginStatus 获取登录状态
//  @param c
func LoginStatus(c *gin.Context) {
	var service service.LoginStatusService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginStatus(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// LoginCellphone 手机登录
//  @param c
func LoginCellphone(c *gin.Context) {
	var service service.LoginCellphoneService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginCellphone(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// LoginEmail 邮箱登录
//  @param c
func LoginEmail(c *gin.Context) {
	var service service.LoginEmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginEmail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// LoginRefresh 刷新登录
//  @param c
func LoginRefresh(c *gin.Context) {
	var service service.LoginRefreshService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginRefresh(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// Logout 退出登录
//  @param c
func Logout(c *gin.Context) {
	var service service.LogoutService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Logout(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
