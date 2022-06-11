package repository

import (
	"dousheng-demo/model"
)

func GetFavoriteVideos(userId int64) []int64 {
	var Ids []int64
	DB.Raw("select video_id from favorites where user_id=?", userId).Scan(&Ids)
	return Ids
}

func SearchFavorite(favorite model.Favorite) bool {
	var t model.Favorite
	t.UserId = "-1"
	DB.Where("video_id = ? ", favorite.VideoId).Where("user_id=?", favorite.UserId).Find(&t)
	if t.UserId == "-1" {
		return false
	} else {
		return true
	}
}
