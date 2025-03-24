package api

import (
	"context"
	"github.com/HCH1212/blog/backend/dao/sqlite"
	"github.com/HCH1212/blog/backend/models"
	"github.com/HCH1212/blog/backend/service"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Token(ctx context.Context, c *app.RequestContext) {
	var err error
	var req models.User
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Fail(c, "参数错误", nil)
		return
	}

	token, refreshToken, err := service.TokenService(sqlite.DB, req)
	if err != nil {
		resp.FailButServer(c, "服务错误", err.Error())
		return
	}
	resp.Success(c, "登陆成功", utils.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	refreshToken := c.PostForm("refresh_token")
	if refreshToken == "" {
		resp.Fail(c, "参数错误", nil)
		return
	}

	token, refreshToken, err := service.RefreshTokenService(refreshToken)
	if err != nil {
		resp.FailButServer(c, "服务错误", err.Error())
		return
	}
	resp.Success(c, "刷新成功", utils.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}
