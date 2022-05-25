package repository

import (
	"dousheng-demo/model"
)

func GetVideo(videoId int64, nextTime int64) model.Video {
	// 从db中获取video
	var video model.Video
	DB.Model(&model.Video{}).Where("create_time < ?", nextTime).Find(&video, videoId)
	return video
}

func GetVideoById(videoId int64) model.Video {
	// 从db中获取video
	var video model.Video
	DB.Model(&model.Video{}).Find(&video, videoId)
	return video
}

//func GetVideoAmount() int64 {
//	// 从db中获取video的数量
//	var count int64
//	DB.Model(&model.Video{}).Where("name <> ?", "").Count(&count)
//	return count
//}

func TimeLimitAmount(timeLimit int64) int64 {
	// 从db中获取截止时间内的数量
	var count int64
	DB.Model(&model.Video{}).Where("create_time < ?", timeLimit).Count(&count)
	return count
}
