package middleware

import (
	"context"
	"github.com/HCH1212/utils/jwt"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
)

// Auth 中间件 - 处理用户身份认证
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Request.Header.Get("Authorization") // 获取 accessToken

		if len(token) == 0 || !strings.HasPrefix(token, "Bearer ") {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
			return
		}

		token = token[7:]
		claims, err := jwt.ParseAccessToken(token)
		if err != nil {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
			return
		}

		//var user *models.User
		//user, err = models.GetUserByID(sqlite.DB, claims.ID)
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	resp.Fail(c, "权限不足", nil)
		//	c.Abort()
		//	return
		//}

		// 用户存在，将信息写入上下文
		c.Set("user_id", claims.ID)
		//c.Set("user", user)
		c.Next(ctx)
	}
}
