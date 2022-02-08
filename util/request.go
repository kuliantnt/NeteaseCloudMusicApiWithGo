package util

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/asmcos/requests"
)

// Options 选项结构
type Options struct {
	Crypto  string
	Ua      string
	Cookies []*http.Cookie
	Token   string
	Url     string
}

// chooseUserAgent 随机选择用户的agent
//  @param ua 浏览器标识
//  @return string
func chooseUserAgent(ua string) string {
	userAgentList := []string{
		//iOS 13.5.1 14.0 beta with safari
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.",
		// iOS with qq micromsg
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/602.1.50 (KHTML like Gecko) Mobile/14A456 QQ/6.5.7.408 V1_IPH_SQ_6.5.7_1_APP_A Pixel/750 Core/UIWebView NetType/4G Mem/103",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.15(0x17000f27) NetType/WIFI Language/zh",
		// Android -> Huawei Xiaomi
		"Mozilla/5.0 (Linux; Android 9; PCT-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.311 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; U; Android 9; zh-cn; Redmi Note 8 Build/PKQ1.190616.001) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/71.0.3578.141 Mobile Safari/537.36 XiaoMi/MiuiBrowser/12.5.22",
		// Android + qq micromsg
		"Mozilla/5.0 (Linux; Android 10; YAL-AL00 Build/HUAWEIYAL-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.62 XWEB/2581 MMWEBSDK/200801 Mobile Safari/537.36 MMWEBID/3027 MicroMessenger/7.0.18.1740(0x27001235) Process/toolsmp WeChat/arm64 NetType/WIFI Language/zh_CN ABI/arm64",
		"Mozilla/5.0 (Linux; U; Android 8.1.0; zh-cn; BKK-AL10 Build/HONORBKK-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/10.6 Mobile Safari/537.36",
		// macOS 10.15.6  Firefox / Chrome / Safari
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:80.0) Gecko/20100101 Firefox/80.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.30 Safari/537.36",
		// Windows 10 Firefox / Chrome / Edge
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.30 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	}

	//设置种子数
	rand.Seed(time.Now().UnixNano())
	index := 0
	if ua == "" {
		index = rand.Intn(len(userAgentList))
	} else if ua == "mobile" {
		index = rand.Intn(8)
	} else {
		index = rand.Intn(5) + 8
	}
	return userAgentList[index]
}

// CreateRequest 创建一个请求
//  @param method 请求方法
//  @param url 请求地址
//  @param data json数据
//  @param options 请求结构
//  @return map[string]interface{} 响应
//  @return []*http.Cookie cookie
func CreateRequest(method string, url string, data map[string]string, options *Options) (map[string]interface{}, []*http.Cookie) {
	req := requests.Requests()
	req.Header.Set("User-Agent", chooseUserAgent(options.Ua))

	// csrfToken 跨站请求伪造保护
	csrfToken := ""

	// music_U todo
	music_U := ""

	music_A := ""

	// answer 响应
	answer := map[string]interface{}{}

	// method 方法大写
	method = strings.ToUpper(method)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if strings.Contains(url, "music.163.com") {
		req.Header.Set("Referer", "https://music.163.com")
	}
	if options.Cookies != nil {
		for _, cookie := range options.Cookies {

			// req 设置cookie
			req.SetCookie(cookie)
			if cookie.Name == "__csrf" {
				csrfToken = cookie.Value
			}
			if cookie.Name == "MUSIC_U" {
				music_U = cookie.Value
			}
			if cookie.Name == "MUSIC_A" {
				music_A = cookie.Value
			}
		}
	}

	if options.Crypto == "weapi" { //如果 交互地址时weapi的话
		data["csrf_token"] = csrfToken
		data = Weapi(data)                   //加密
		reg, _ := regexp.Compile(`/\w*api/`) //将所有xxxapi替换为weapi
		url = reg.ReplaceAllString(url, "/weapi/")
	} else if options.Crypto == "linuxapi" { //如果 加密方法为linuxapi的话
		linuxApiData := make(map[string]interface{}, 3)
		linuxApiData["method"] = method
		reg, _ := regexp.Compile(`/\w*api/`)
		linuxApiData["url"] = reg.ReplaceAllString(url, "/api/") //将所有xxxapi替换为weapi
		linuxApiData["params"] = data
		data = Linuxapi(linuxApiData) //加密
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
		url = "https://music.163.com/api/linux/forward"
	} else if options.Crypto == "eapi" {
		eapiData := make(map[string]interface{})
		for key, value := range data {
			eapiData[key] = value
		}
		rand.Seed(time.Now().UnixNano())
		header := map[string]string{
			"osver":       "",      //系统版本
			"deviceId":    "",      //encrypt.base64.encode(imei + '\t02:00:00:00:00:00\t5106025eb79a5247\t70ffbaac7')
			"versioncode": "140",   //版本号
			"appver":      "8.0.0", // app版本
			"mobilename":  "",      //设备model
			"buildver":    strconv.FormatInt(time.Now().Unix(), 10),
			"resolution":  "1920x1080", //设备分辨率
			"os":          "android",
			"channel":     "",
			"requestId":   strconv.FormatInt(time.Now().Unix()*1000, 10) + strconv.Itoa(rand.Intn(1000)),
			"MUSIC_U":     music_U,
			"MUSIC_A":     music_A,
		}

		for key, value := range header {
			req.SetCookie(&http.Cookie{Name: key, Value: value, Path: "/"})
		}
		eapiData["header"] = header
		data = Eapi(options.Url, eapiData)
		reg, _ := regexp.Compile(`/\w*api/`)
		url = reg.ReplaceAllString(url, "/eapi/")
	}
	var resp *requests.Response
	var err error
	if method == "POST" {
		var form requests.Datas = data
		resp, err = req.Post(url, form)
	} else {
		resp, err = req.Get(url)
	}

	if err != nil {
		answer["code"] = 520
		answer["err"] = err.Error()
		return answer, nil
	}
	cookies := resp.Cookies()

	body := resp.Content()
	//fmt.Println(string(body))
	b := bytes.NewReader(body)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	// 数据被压缩 进行解码
	if err == nil {
		io.Copy(&out, r)
		body = out.Bytes()
	}

	err = json.Unmarshal(body, &answer)
	// 出错说明不是json
	if err != nil {
		//fmt.Println(string(body))
		// 可能是纯页面
		if strings.Contains(string(body), "<!DOCTYPE html>") {
			answer["code"] = 200
			answer["html"] = string(body)
			return answer, cookies
		}
		answer["code"] = 500
		answer["err"] = err.Error()
		fmt.Println(string(body))
		return answer, nil
	}
	if _, ok := answer["code"]; !ok {
		answer["code"] = 200
	}
	return answer, cookies
}
