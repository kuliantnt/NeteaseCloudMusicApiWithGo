package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type DjProgramService struct {
	RID    string `json:"rid" form:"rid"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
	Asc    string `json:"asc" form:"asc"`
}

func (service *DjProgramService) DjProgram(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["radioId"] = service.RID
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["asc"] = service.Asc
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/dj/program/byradio`, data, options)

	return reBody
}
