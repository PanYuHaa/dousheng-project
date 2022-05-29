package service

import "dousheng-demo/repository"

func InfoVerify(enterpassword string, enterusername string, id int64) bool {
	if enterusername == repository.GetUsernameById(id) && enterpassword == repository.GetPasswordById(id) {
		return true
	} else {
		return false
	}
}
