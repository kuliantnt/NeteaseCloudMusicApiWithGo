package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type TopPlaylistService struct {
	Limit  string `json:"limit" form:"limit"`
	Cat    string `json:"cat" form:"cat"`
	Order  string `json:"order" form:"order"`
	Offset string `json:"offset" form:"offset"`
}

func (service *TopPlaylistService) TopPlaylist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Cat == "" {
		service.Cat = "全部"
	}
	if service.Limit == "" {
		service.Limit = "50"
	}
	if service.Offset == "" {
		service.Offset = "0"
	}
	if service.Order != "hot" {
		service.Order = "new"
	}
	data["limit"] = service.Limit
	data["offset"] = service.Offset
	data["total"] = "true"
	data["hot"] = service.Order
	data["cat"] = service.Cat

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/list`, data, options)

	return reBody
}
