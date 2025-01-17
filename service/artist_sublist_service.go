package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type ArtistSublistService struct {
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *ArtistSublistService) ArtistSublist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Limit == "" {
		service.Limit = "25"
	}
	if service.Offset == "" {
		service.Offset = "0"
	}
	data["limit"] = service.Limit
	data["offset"] = service.Offset
	data["total"] = "true"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/artist/sublist`, data, options)

	return reBody
}
