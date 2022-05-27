package service

import (
	"dousheng-demo/repository"
)

func GetNextUserId() int64 {
	return repository.GetUsersAmount()
}

func IsUserExist(username string) bool {
	return repository.IsUserExist(username)
}

//func RegisterAccount(userClaim model.UserClaim) error {
//	return repository.AddUserClaim(userClaim)
//}
