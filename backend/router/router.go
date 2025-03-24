package router

import (
	"context"
	"github.com/HCH1212/blog/backend/api"
	"github.com/HCH1212/blog/backend/middleware"
	umiddleware "github.com/HCH1212/utils/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net/http"
)

func Init() {
	h := server.Default()
	h.Use(umiddleware.Cors) // 跨域中间件

	// 用户相关
	u := h.Group("/user")
	{
		u.POST("/token", api.Token)                // 用户登陆
		u.POST("/refresh_token", api.RefreshToken) // 刷新token
		//u.GET("/info", middleware.Auth(), func(ctx context.Context, c *app.RequestContext) {
		//	res, _ := c.Get("user_id")
		//	c.JSON(http.StatusOK, res.(uint))
		//})
	}

	// 文章相关
	a := h.Group("/article")
	{
		a.POST("/add", middleware.Admin(), api.AddArticle)           // 管理员添加文章
		a.POST("/update", middleware.Admin(), api.UpdateArticle)     // 管理员修改文章
		a.POST("/delete", middleware.Admin(), api.DeleteArticle)     // 管理员删除文章
		a.GET("/list", api.ListArticles)                             // 文章列表
		a.GET("/get", api.GetArticle)                                // 文章详情
		a.POST("/favorite", middleware.Auth(), api.AddFavorite)      // 用户添加收藏
		a.DELETE("/favorite", middleware.Auth(), api.RemoveFavorite) // 用户取消收藏
		a.GET("/search", api.Search)                                 // 搜索文章
		a.GET("/list/tag", api.GetArticlesByTag)                     // 根据tag分类文章
		a.GET("/tags", api.GetTags)                                  // 获取所有标签
	}

	// 评论相关
	c := h.Group("/comment")
	{
		c.POST("/add", middleware.Auth(), api.AddComment)       // 用户新增评论
		c.POST("/delete", middleware.Auth(), api.DeleteComment) // 用户删除评论
		c.GET("/list", api.ListComments)                        // 列出评论
		c.GET("/list/son", api.GetReplies)                      // 列出子评论
	}

	// 图床相关
	i := h.Group("/image")
	{
		i.POST("/upload", middleware.Auth(), middleware.UploadLimit(), api.Upload) // 图片上传，每个用户只能上传12次
	}

	// ai相关
	ai := h.Group("/ai")
	{
		ai.POST("/chat", api.Chat) // ai聊天
	}

	// test
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(http.StatusOK, "pong")
	})
	h.Spin()
}
