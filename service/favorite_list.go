package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
	"fmt"
	"strconv"
)

var u = 0

func FavoriteList(UserId string) []model.Video {
	VideoList := make([]model.Video, 0)
	VideoIds := repository.GetFavoriteVideos(UserId)
	for _, Id := range VideoIds {
		u++
		Id64, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			fmt.Printf("wrong!")
		}
		VideoList = append(VideoList, repository.GetVideoById(Id64))
	}

	return VideoList
}

func FavoriteListRsp(UserId string) bool {
	if u != 0 {
		return true
	} else {
		return false
	}
}
