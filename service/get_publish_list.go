package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetPublishList(userId int64) []model.Video {
	userName := repository.GetUserById(userId).Name
	return repository.GetVideosByName(userName)
}
