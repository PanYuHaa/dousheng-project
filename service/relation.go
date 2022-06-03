package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func Follow(userid string, toid string) error {
	newSubscribe := model.Follow{UserId: userid,
		ToUserId: toid}
	return repository.AddNewFollow(newSubscribe)
}

func UnFollow(userid string, toid string) error {
	newSubscribe := model.Follow{UserId: userid,
		ToUserId: toid}
	return repository.DeleteFollow(newSubscribe)
}

func SearchFollow(userid string, toid string) bool {
	newSubscribe := model.Follow{UserId: userid,
		ToUserId: toid}
	return repository.SearchFollow(newSubscribe)
}
