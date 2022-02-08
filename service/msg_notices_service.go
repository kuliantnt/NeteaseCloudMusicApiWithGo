package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type MsgNoticesService struct {
	Limit    string `json:"limit" form:"limit"`
	LastTime string `json:"lasttime" form:"lasttime"`
}

func (service *MsgNoticesService) MsgNotices(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.LastTime == "" {
		data["time"] = "-1"
	} else {
		data["time"] = service.LastTime
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/msg/notices`, data, options)

	return reBody
}
