package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type VideoDetailInfoService struct {
	ID string `json:"vid" form:"vid"`
}

func (service *VideoDetailInfoService) VideoDetailInfo(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["threadid"] = "R_VI_62_" + service.ID
	data["composeliked"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/comment/commentthread/info`, data, options)

	return reBody
}
