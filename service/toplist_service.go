package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type ToplistService struct {
}

func (service *ToplistService) Toplist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "linuxapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/toplist`, data, options)

	return reBody
}
