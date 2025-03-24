package service

import (
	"errors"
	"github.com/HCH1212/blog/backend/models"

	"github.com/HCH1212/utils/jwt"
	"github.com/HCH1212/utils/password"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func TokenService(db *gorm.DB, u models.User) (string, string, error) {
	dataUSer, err := models.GetUserByName(db, u.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 执行注册逻辑
			var hashPassword string
			hashPassword, err = password.HashPassword(u.Password)
			if err != nil {
				logrus.Info(err)
				return "", "", err
			}
			newUser := &models.User{Username: u.Username, Password: hashPassword}
			err = models.CreateUser(db, newUser)
			if err != nil {
				logrus.Info(err)
				return "", "", err
			}
			return jwt.GetToken(newUser.ID)
		} else {
			logrus.Info(err)
			return "", "", err
		}
	}
	// 登陆逻辑
	if !password.VerifyPassword(dataUSer.Password, u.Password) {
		return "", "", errors.New("password error")
	}
	return jwt.GetToken(dataUSer.ID)
}

func RefreshTokenService(refreshToken string) (string, string, error) {
	res, err := jwt.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	return jwt.GetToken(res.ID)
}
