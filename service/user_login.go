package service

import (
	"dousheng-demo/repository"
	"golang.org/x/crypto/bcrypt"
)

func IdentityVerify(username string, password string) bool {
	account := repository.GetAccount(username)
	if account.UserName == username {
		err := bcrypt.CompareHashAndPassword([]byte(account.PassWord), []byte(password)) //验证（对比）
		if err != nil {
			return false
		}
		return true
	}
	return false
}
