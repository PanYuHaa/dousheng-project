package service

import "dousheng-demo/repository"

func ComparePassword(enterpassword string, username string) bool {
	turepassword := repository.GetPasswordByUsername(username)
	if turepassword == enterpassword {
		return true
	} else {
		return false
	}
}
func GetUseridByName(username string) int64 {
	return repository.GetUseridByName(username)
}
