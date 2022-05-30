package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func Favorite(userid string, videoid string) error {
	newfavorite := model.Favorite{UserId: userid,
		VideoId: videoid}
	return repository.AddNewFavorite(newfavorite)
}
func Disfavorite(userid string, videoid string) error {
	newfavorite := model.Favorite{UserId: userid,
		VideoId: videoid}
	return repository.DeleteFavorite(newfavorite)
}
