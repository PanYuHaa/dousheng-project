package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

var u = 0

func FavoriteList(userId int64) []model.Video {
	VideoList := make([]model.Video, 0)
	VideoIds := repository.GetFavoriteVideos(userId)
	for _, id := range VideoIds {
		u++
		VideoList = append(VideoList, repository.GetVideoById(id))
	}

	return VideoList
}

func FavoriteListRsp() bool {
	if u != 0 {
		return true
	} else {
		return false
	}
}
