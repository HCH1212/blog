package middleware

import (
	"context"
	"errors"
	"github.com/HCH1212/blog/backend/dao/sqlite"
	"github.com/HCH1212/blog/backend/models"
	"github.com/HCH1212/utils/jwt"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
	"strings"
)

// Admin 中间件 - 处理用户权限
func Admin() app.HandlerFunc {
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

		var user *models.User
		user, err = models.GetUserByID(sqlite.DB, claims.ID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
			return
		}

		if user == nil {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
			return
		}

		if user.Role == 1 {
			c.Next(ctx)
		} else {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
		}
	}
}
