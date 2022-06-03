package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func AddUser(user model.User, account model.Account) error {
	return repository.AddUser(user, account)
}
