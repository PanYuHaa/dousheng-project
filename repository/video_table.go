package repository

import (
	"dousheng-demo/model"
)

//	Add

func AddVideo(video model.Video) error {
	dbRes := DB.Model(&model.Video{}).Create(&video)
	return dbRes.Error
}
func AddVideoFavorite(VideoId string) error {
	mu.Lock()
	defer mu.Unlock()
	var video model.Video
	dbRes := DB.First(&video, VideoId)
	video.FavoriteCount++
	DB.Save(&video)
	return dbRes.Error
}
func DeleteVideoFavorite(VideoId string) error {
	mu.Lock()
	defer mu.Unlock()
	var video model.Video
	dbRes := DB.First(&video, VideoId)
	video.FavoriteCount--
	DB.Save(&video)
	return dbRes.Error
}

//	Get

func GetVideoById(videoId int64) model.Video {
	// 从db中获取video
	var video model.Video
	DB.Model(&model.Video{}).Find(&video, videoId)
	return video
}

func GetVideosByName(nickname string) []model.Video {
	var videos []model.Video
	DB.Model(&model.Video{}).Where("name = ?", nickname).Find(&videos)
	return videos
}

//	Other

func TimeLimitAmount(timeLimit int64) int64 {
	// 从db中获取截止时间内的数量
	var count int64
	DB.Model(&model.Video{}).Where("create_time < ?", timeLimit).Count(&count)
	return count
}

func VideoAmount() int64 {
	// 从db中获取截止时间内的数量
	var count int64
	DB.Model(&model.Video{}).Where("name != ?", "").Count(&count)
	return count
}
