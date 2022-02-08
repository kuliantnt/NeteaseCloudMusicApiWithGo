package main

import (
	"github.com/lianlian/singo/conf"
	"github.com/lianlian/singo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3333")
}
