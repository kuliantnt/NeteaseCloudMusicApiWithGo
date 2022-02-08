package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type DjToplistNewcomerService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *DjToplistNewcomerService) DjToplistNewcomer(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "100"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/dj/toplist/newcomer`, data, options)

	return reBody
}
