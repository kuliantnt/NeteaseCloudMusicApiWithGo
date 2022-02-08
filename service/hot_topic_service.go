package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type HotTopicService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *HotTopicService) HotTopic(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "20"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/act/hot`, data, options)

	return reBody
}
