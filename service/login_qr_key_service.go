package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lianlian/singo/util"
)

// LoginQrKeyService 二维码
type LoginQrKeyService struct{}

// LoginQrkey 二维码
//  @receiver service
//  @param c
//  @return map
func (service *LoginQrKeyService) LoginQrkey(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	//传递的数据
	data := make(map[string]interface{})
	data["type"] = 1

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/login/qrcode/unike`, data, options)

	return reBody
}
