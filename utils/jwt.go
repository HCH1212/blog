package utils

import (
	"blog/global"
	"blog/model"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// TODO 这里初始化jwt key会因为配置没有初始化好而造成空指针
//var accessTokenKey = []byte(global.Config.JWT.AccessKey)
//var refreshTokenKey = []byte(global.Config.JWT.RefreshKey)

type Claims struct {
	ID uint
	jwt.RegisteredClaims
}

// 使用双token
func GetToken(user *model.User) (string, string, error) {
	var accessTokenKey = []byte(global.Config.JWT.AccessKey)
	var refreshTokenKey = []byte(global.Config.JWT.RefreshKey)

	// accessToken过期时间一周, refreshToken过期时间一月
	accessTokenTime := time.Now().Add(7 * 24 * time.Hour)
	refreshTokenTime := time.Now().Add(4 * 7 * 24 * time.Hour)

	accessClaims := Claims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my",
			Subject:   "token",
		},
	}
	refreshClaims := Claims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my",
			Subject:   "token",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenStr, err := accessToken.SignedString(accessTokenKey)
	if err != nil {
		return "", "", err
	}
	refreshTokenStr, err := refreshToken.SignedString(refreshTokenKey)
	if err != nil {
		return "", "", err
	}
	return accessTokenStr, refreshTokenStr, nil
}

func ParseAccessToken(tokenString string) (*Claims, error) {
	var accessTokenKey = []byte(global.Config.JWT.AccessKey)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return accessTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func ParseRefreshToken(tokenString string) (*Claims, error) {
	var refreshTokenKey = []byte(global.Config.JWT.RefreshKey)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
