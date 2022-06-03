package service

import "dousheng-demo/repository"

func IdentityVerify(username string, password string) bool {
	account := repository.GetAccount(username, password)
	if account.UserName == username && account.PassWord == password {
		return true
	}
	return false
}
