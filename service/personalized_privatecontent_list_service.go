package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type PersonalizedPrivatecontentListService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *PersonalizedPrivatecontentListService) PersonalizedPrivatecontentList(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "60"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v2/privatecontent/list`, data, options)

	return reBody
}
