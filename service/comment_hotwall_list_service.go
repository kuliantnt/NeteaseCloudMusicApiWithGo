package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type CommentHotwallListService struct {
}

func (service *CommentHotwallListService) CommentHotwallList(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/comment/hotwall/list/get`, data, options)

	return reBody
}
