package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func AddUser(user model.User) error {
	return repository.AddUser(user)
}
