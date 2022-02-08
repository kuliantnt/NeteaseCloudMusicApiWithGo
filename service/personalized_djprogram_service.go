package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type PersonalizedDjprogramService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *PersonalizedDjprogramService) PersonalizedDjprogram(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/djprogram`, data, options)

	return reBody
}
