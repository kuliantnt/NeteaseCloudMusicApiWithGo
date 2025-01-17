package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type SimiArtistService struct {
	ID string `json:"id" form:"id"`
}

func (service *SimiArtistService) SimiArtist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/discovery/simiArtist`, data, options)

	return reBody
}
