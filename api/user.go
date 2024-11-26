package api

import (
	"blog/model"
	"blog/resp"
	"blog/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		resp.Fail(c, "传入参数格式错误", nil)
		return
	}
	if len(user.Telephone) == 0 || len(user.Telephone) != 11 || len(user.Password) == 0 {
		resp.Fail(c, "传入参数内容错误", nil)
		return
	}
	err = service.RegisterService(&user)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "注册成功", nil)
}

func Login(c *gin.Context) {
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	if len(telephone) == 0 || len(telephone) != 11 || len(password) == 0 {
		resp.Fail(c, "传入参数内容错误", nil)
		return
	}
	accessToken, refreshToken, err := service.LoginService(telephone, password)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "登陆成功", gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")
	if refreshToken == "" {
		resp.Fail(c, "传入参数错误", nil)
		return
	}
	accessToken, refreshToken, err := service.RefreshTokenService(refreshToken)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "刷新Token成功", gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}
