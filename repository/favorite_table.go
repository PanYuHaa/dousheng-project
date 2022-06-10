package repository

import (
	"dousheng-demo/model"
)

//func AddNewFavorite(favorite model.Favorite) error {
//	mu.Lock()
//	defer mu.Unlock()
//	//dbRes := DB.Model(&model.Favorite{}).Create(&favorite)
//	return DB.Transaction(func(tx *gorm.DB) error {
//		// favorite table
//		if err := tx.Model(&model.Favorite{}).Create(&favorite).Error; err != nil {
//			return err
//		}
//		// videos table
//		if err := tx.Table("videos").Where("user_id = ?", favorite.ToUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
//			return err
//		}
//		if err := tx.Table("videos").Where("user_id = ?", subscribe.UserId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
//			return err
//		}
//		return nil
//	})
//}

//func DeleteFavorite(favorite model.Favorite) error {
//	mu.Lock()
//	defer mu.Unlock()
//	dbRes := DB.Where("video_id = ? ", favorite.VideoId).Where("user_id = ?", favorite.UserId).Delete(&model.Favorite{})
//	return dbRes.Error
//}

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
