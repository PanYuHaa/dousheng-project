package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func Favorite(userid string, videoid string) error {
	newfavorite := model.Favorite{UserId: userid,
		VideoId: videoid}
	t := repository.AddVideoFavorite(videoid)
	if t == nil {
		return repository.AddNewFavorite(newfavorite)
	} else {
		return t
	}

}
func Disfavorite(userid string, videoid string) error {
	newfavorite := model.Favorite{UserId: userid,
		VideoId: videoid}
	t := repository.DeleteVideoFavorite(videoid)
	if t == nil {
		return repository.DeleteFavorite(newfavorite)
	} else {
		return t
	}
}
func SearchFavorite(userid string, videoid string) bool {
	newfavorite := model.Favorite{UserId: userid,
		VideoId: videoid}
	return repository.SearchFavorite(newfavorite)
}
