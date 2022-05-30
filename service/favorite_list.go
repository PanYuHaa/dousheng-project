package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
	"fmt"
	"strconv"
)

var u = 0

func FavoriteListRsp(UserId string) model.Response {
	if u != 0 {
		return model.Response{StatusCode: 1, StatusMsg: "Favorite Video"}
	} else {
		return model.Response{StatusCode: -1, StatusMsg: "No Favorite Videos"}
	}
}
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
