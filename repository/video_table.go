package repository

import (
	"dousheng-demo/model"
	"gorm.io/gorm"
)

//	Add

func AddVideo(video model.Video) error {
	dbRes := DB.Model(&model.Video{}).Create(&video)
	return dbRes.Error
}
func AddVideoFavorite(VideoId string, favorite model.Favorite) error {
	mu.Lock()
	defer mu.Unlock()
	var video model.Video
	//dbRes := DB.First(&video, VideoId)
	//video.FavoriteCount++
	//DB.Save(&video)
	//return dbRes.Error
	return DB.Transaction(func(tx *gorm.DB) error {
		// favorite table
		if err := tx.Model(&model.Favorite{}).Create(&favorite).Error; err != nil {
			return err
		}
		// videos table
		if err := tx.First(&video, VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}
func DeleteVideoFavorite(VideoId string, favorite model.Favorite) error {
	mu.Lock()
	defer mu.Unlock()
	var video model.Video
	//dbRes := DB.First(&video, VideoId)
	//video.FavoriteCount--
	//DB.Save(&video)
	//return dbRes.Error
	return DB.Transaction(func(tx *gorm.DB) error {
		// favorite table
		if err := tx.Model(&model.Favorite{}).Where("user_id = ? ", favorite.UserId).Where("video_id = ?", favorite.VideoId).Delete(&favorite).Error; err != nil {
			return err
		}
		// videos table
		if err := tx.First(&video, VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", -1)).Error; err != nil {
			return err
		}
		return nil
	})
}

//	Get

func GetVideoById(videoId int64) model.Video {
	// 从db中获取video
	var video model.Video
	DB.Model(&model.Video{}).Find(&video, videoId)
	return video
}

func GetVideosByUserId(user model.User) []model.Video {
	var videos []model.Video
	DB.Table("videos").Where("user_id = ?", user.UserId).Find(&videos)
	return videos
}

//	Other

func TimeLimitAmount(timeLimit int64) int64 {
	// 从db中获取截止时间内的数量
	var count int64
	DB.Model(&model.Video{}).Where("create_time < ?", timeLimit).Count(&count)
	return count
}

//func VideoAmount() int64 {
//	// 从db中获取视频的数量
//	var count int64
//	DB.Model(&model.Video{}).Where("name != ?", "").Count(&count)
//	return count
//}
