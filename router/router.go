package router

import (
	"blog/api"
	"blog/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	i := router.Group("image")
	{
		i.POST("/update", api.ImageUpdate)
	}

	addr := global.Config.System.Addr()
	global.Log.Infof("blog server run at %s", addr)
	_ = router.Run(addr)
}
