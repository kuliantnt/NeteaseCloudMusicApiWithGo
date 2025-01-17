package service

import (
	"net/http"

	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type CommentHotService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
	Before string `json:"before" form:"before"`
	Type   string `json:"type" form:"type"`
}

func (service *CommentHotService) CommentHot(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	TYPE := make(map[string]string, 6)
	TYPE["0"] = "R_SO_4_"
	TYPE["1"] = "R_MV_5_"
	TYPE["2"] = "A_PL_0_"
	TYPE["3"] = "R_AL_3_"
	TYPE["4"] = "A_DJ_1_"
	TYPE["5"] = "R_VI_62_"
	data := make(map[string]string)

	data["rid"] = service.ID
	if _, ok := TYPE[service.Type]; ok {
		service.Type = TYPE[service.Type]
	} else {
		service.Type = TYPE["0"]
	}
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
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/resource/hotcomments/`+service.Type+service.ID, data, options)

	return reBody
}
