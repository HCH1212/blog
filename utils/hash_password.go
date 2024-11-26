package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func EncryptPassword(password string) (string, error) {
	saltedPassword := password
	bytes, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证密码
func VerifyPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
