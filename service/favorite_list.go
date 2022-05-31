package service

import (
	"dousheng-demo/controller"
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
func FavoriteListRsp(UserId string) controller.Response {
	if u != 0 {
		return controller.Response{StatusCode: 1, StatusMsg: "Favorite Video"}
	} else {
		return controller.Response{StatusCode: -1, StatusMsg: "No Favorite Videos"}
	}
}
