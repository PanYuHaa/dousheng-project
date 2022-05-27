package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetPublishList(userId int64) []model.Video {
	nickname := repository.GetUserById(userId).NickName
	return repository.GetVideosByName(nickname)
}
