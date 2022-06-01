package repository

import (
	"dousheng-demo/model"
)

func AddNewFavorite(favorite model.Favorite) error {
	dbRes := DB.Model(&model.Favorite{}).Create(&favorite)
	return dbRes.Error
}
func DeleteFavorite(favorite model.Favorite) error {
	dbRes := DB.Where("video_id = ? ", favorite.VideoId).Where("user_id = ?", favorite.UserId).Delete(&model.Favorite{})
	return dbRes.Error
}

func GetFavoriteVideos(UserId string) []string {
	var Ids []string
	DB.Raw("select video_id from favorites where user_id=?", UserId).Scan(&Ids)
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
