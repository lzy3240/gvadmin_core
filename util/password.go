package util

import (
	"golang.org/x/crypto/bcrypt"
	"gvadmin_v3/core/global/E"
)

// 密码加密
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd+E.Salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// 密码验证
func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd+E.Salt))

	return err == nil
}
