package repository

import "dousheng-demo/model"

func AddNewFavorite(favorite model.Favorite) error {
	dbRes := DB.Model(&model.Favorite{}).Create(&favorite)
	return dbRes.Error
}
func DeleteFavorite(favorite model.Favorite) error {
	dbRes := DB.Where("video_id = ? ", favorite.VideoId).Where("user_id=?", favorite.UserId).Delete(&model.Favorite{})
	return dbRes.Error
}

//func GetFavoriteVideo()model.Video{
//
//}
