package service

import (
	"net/http"

	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type CommentDjService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
	Before string `json:"before" form:"before"`
}

func (service *CommentDjService) CommentDj(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["rid"] = service.ID
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
	if service.Before == "" {
		data["beforeTime"] = "0"
	} else {
		data["beforeTime"] = service.Before
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/resource/comments/A_DJ_1_`+service.ID, data, options)

	return reBody
}
