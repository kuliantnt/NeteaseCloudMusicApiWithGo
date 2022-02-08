package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type DjToplistPopularService struct {
	Limit string `json:"limit" form:"limit"`
}

func (service *DjToplistPopularService) DjToplistPopular(c *gin.Context) map[string]interface{} {

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

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/dj/toplist/popular`, data, options)

	return reBody
}
