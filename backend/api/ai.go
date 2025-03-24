package api

import (
	"context"
	"github.com/HCH1212/blog/backend/service"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

func Chat(ctx context.Context, c *app.RequestContext) {
	content := c.PostForm("content")

	res, err := service.ChatService(ctx, content)
	if err != nil {
		resp.FailButServer(c, "服务错误", err.Error())
		return
	}

	resp.Success(c, "请求成功", res)
}
