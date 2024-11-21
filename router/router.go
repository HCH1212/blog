package router

import (
	"github.com/HCH1212/blog/blog_server/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	settingRouter(router)

	addr := global.Config.System.Addr()
	global.Log.Infof("blog server run at %s", addr)
	_ = router.Run(addr)
}
