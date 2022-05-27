package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetNextUserId() int64 {
	return repository.GetUsersAmount()
}

func IsUserExist(username string) bool {
	return repository.IsUserExist(username)
}

func AddAccount(account model.Account) error {
	return repository.AddAccount(account)
}
