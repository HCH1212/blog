package main

import (
	"github.com/HCH1212/blog/blog_server/core"
	"github.com/HCH1212/blog/blog_server/router"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	core.InitLogger()
	// 连接数据库
	core.InitGorm()
	// 路由
	router.InitRouter()
}
