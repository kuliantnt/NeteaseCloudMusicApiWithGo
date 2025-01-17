package service

import (
	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type AlbumSongsaleboardService struct {
	AlbumType string `json:"albumType" form:"albumType"`
	Limit     string `json:"limit" form:"limit"`
	Offset    string `json:"offset" form:"offset"`
	Type      string `json:"type" form:"type"`
	Year      string `json:"year" form:"year"`
}

func (service *AlbumSongsaleboardService) AlbumSongsaleboard(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	if service.AlbumType == "" {
		service.AlbumType = "0"
	}
	if service.Type == "" {
		service.Type = "daily"
	}
	data["albumType"] = service.AlbumType
	if service.Type == "year" {
		data["year"] = service.Year
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/feealbum/songsaleboard/`+service.Type+"/type", data, options)

	return reBody
}
