package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetPublishList(userId int64) []model.Video {
	user := repository.GetUserById(userId)
	return repository.GetVideosByUserId(user)
}
