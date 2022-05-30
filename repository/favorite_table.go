package repository

import (
	"dousheng-demo/model"
	"gorm.io/gorm"
)

func AddNewFavorite(favorite model.Favorite) error {
	dbRes := DB.Model(&model.Favorite{}).Create(&favorite)
	return dbRes.Error
}
func DeleteFavorite(favorite model.Favorite) error {
	dbRes := DB.Where("video_id = ? ", favorite.VideoId).Where("user_id=?", favorite.UserId).Delete(&model.Favorite{})
	return dbRes.Error
}

func GetFavoriteVideos(UserId string) []string {
	var Ids []string
	DB.Where("user_id = ?", UserId).FindInBatches(&Ids, 100, func(tx *gorm.DB, batch int) error {
		// 如果返回错误会终止后续批量操作
		return nil
	})
	return Ids
}
