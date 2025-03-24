package middleware

import (
	"context"
	"fmt"
	"github.com/HCH1212/blog/backend/dao/redis"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

// UploadLimit 限制每个用户的上传次数
func UploadLimit() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从上下文中获取用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			resp.Fail(c, "无法上传", nil)
			c.Abort()
			return
		}

		// 构造Redis键
		key := fmt.Sprintf("upload_limit:%s", userID)

		// 获取当前上传次数
		count, err := redis.ReClient.Get(ctx, key).Int()
		if err != nil {
			if err.Error() == "redis: nil" {
				// 第一次
				redis.ReClient.Set(ctx, key, 0, 0)
			} else {
				resp.Fail(c, "无法上传", nil)
				c.Abort()
				return
			}
		}

		// 如果上传次数超过限制，返回错误
		if count > 12 {
			resp.Fail(c, "无法上传", nil)
			c.Abort()
			return
		}

		// 增加上传次数
		err = redis.ReClient.Incr(ctx, key).Err()
		if err != nil {
			resp.Fail(c, "无法上传", nil)
			c.Abort()
			return
		}

		// 继续处理请求
		c.Next(ctx)
	}
}
