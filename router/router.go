package router

import (
	"blog/api"
	"blog/global"
	"blog/middleware"
	"blog/resp"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "111")
	})

	r := router.Group("/user")
	{
		r.POST("/register", api.Register)                        // 注册
		r.POST("/login", api.Login)                              // 登陆
		r.POST("/refresh_token", api.RefreshToken)               // refreshToken刷新token
		r.GET("/info", middleware.Auth(), func(c *gin.Context) { // 验证中间件
			user, _ := c.Get("user")
			resp.Success(c, "成功验证token并获取用户信息", user)
		})
	}

	addr := global.Config.System.Addr()
	global.Log.Infof("blog server run at %s", addr)
	_ = router.Run(addr)
}
