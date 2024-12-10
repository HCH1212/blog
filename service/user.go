package service

import (
	"blog/global"
	"blog/model"
	"blog/utils"
	"errors"
)

func RegisterService(user *model.User) error {
	hashPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return errors.New("服务错误")
	}
	user.Password = hashPassword
	res := global.DB.Where("telephone = ?", user.Telephone).FirstOrCreate(user)
	if res.Error != nil {
		global.Log.Error(res.Error.Error())
		return errors.New("服务错误")
	}
	if res.RowsAffected != 1 {
		return errors.New("用户已存在")
	}
	return nil
}

func LoginService(telephone, password string) (string, string, error) {
	var user model.User
	res := global.DB.Where("telephone = ?", telephone).First(&user)
	if res.Error != nil {
		global.Log.Error(res.Error.Error())
		return "", "", errors.New("服务错误")
	}
	if !utils.VerifyPassword(user.Password, password) {
		return "", "", errors.New("密码错误")
	}
	accessToken, refreshToken, err := utils.GetToken(&user)
	if err != nil {
		global.Log.Error(err.Error())
		return "", "", errors.New("服务错误")
	}
	return accessToken, refreshToken, nil
}

func RefreshTokenService(refreshToken string) (string, string, error) {
	res, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		global.Log.Error(err.Error())
		return "", "", errors.New("服务错误")
	}
	var user model.User
	user.ID = res.ID
	accessToken, refreshToken, err := utils.GetToken(&user)
	if err != nil {
		global.Log.Error(err.Error())
		return "", "", errors.New("服务错误")
	}
	return accessToken, refreshToken, nil
}
