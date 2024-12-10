package middleware

import (
	"blog/global"
	"blog/model"
	"blog/resp"
	"blog/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization") // 传入accessToken

		if len(token) == 0 || !strings.HasPrefix(token, "Bearer ") {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
			return
		}

		token = token[7:]
		claims, err := utils.ParseAccessToken(token)
		if err != nil {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
			return
		}

		id := claims.ID
		var user model.User
		res := global.DB.Find(&user, id)
		// 用户已经不存在
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			resp.Fail(c, "权限不足", nil)
			c.Abort()
			return
		}
		// 用户存在，将信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
