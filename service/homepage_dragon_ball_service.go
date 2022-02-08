package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type HomepageDragonBallService struct {
}

func (service *HomepageDragonBallService) HomepageDragonBall(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/homepage/dragon/ball/static",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/eapi/homepage/dragon/ball/static`, data, options)

	return reBody
}
