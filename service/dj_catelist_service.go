package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type DjCatelistService struct {
}

func (service *DjCatelistService) DjCatelist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/category/get`, data, options)

	return reBody
}
