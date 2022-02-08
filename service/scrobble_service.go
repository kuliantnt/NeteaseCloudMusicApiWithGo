package service

import (
	"encoding/json"

	"github.com/lianlian/singo/util"

	"github.com/gin-gonic/gin"
)

type ScrobbleService struct {
	ID       string `json:"id" form:"id"`
	Sourceid string `json:"sourceid" form:"sourceid"`
	Time     string `json:"time" form:"time"`
}

func (service *ScrobbleService) Scrobble(c *gin.Context) map[string]interface{} {

	//errBody:=make(map[string]interface{})
	//errBody["code"]=500
	//errBody["err"]="此接口本后台暂未实现，，，欢迎pr"
	//return errBody

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	jsonn := make(map[string]interface{})
	jsonn["download"] = 0
	jsonn["end"] = "playend"
	jsonn["id"] = service.ID
	jsonn["sourceId"] = service.Sourceid
	jsonn["time"] = service.Time
	jsonn["type"] = "song"
	jsonn["wifi"] = 0

	long := make(map[string]interface{})
	long["action"] = "play"
	long["json"] = jsonn

	var longs []map[string]interface{}
	longs = append(longs, long)

	if str, err := json.Marshal(longs); err != nil {
		errBody := make(map[string]interface{})
		errBody["code"] = 502
		errBody["err"] = "参数错误"
		return errBody
	} else {
		data["long"] = string(str)
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/feedback/weblog`, data, options)

	return reBody
}
