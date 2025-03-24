package main

import (
	"github.com/HCH1212/blog/backend/dao"
	"github.com/HCH1212/blog/backend/router"
	"github.com/HCH1212/utils/log"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	log.InitDefaultLogger("[blog]") // 日志初始化
	dao.Init()
	router.Init()

}
