package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetPublishList(userId int64) []model.Video {
	name := repository.GetUserById(userId).Name
	return repository.GetVideosByName(name)
}
