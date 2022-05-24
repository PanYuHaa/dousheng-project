package repository

import (
	"dousheng-demo/model"
)

func GetVideo(videoId int64) model.Video {
	// 从db中获取video
	var TestVideo model.Video
	DB.Model(&model.Video{}).Find(&TestVideo, videoId)
	return TestVideo
}

func GetVideoAmount() int64 {
	// 从db中获取video的数量
	var count int64
	DB.Model(&model.Video{}).Where("name <> ?", "").Count(&count)
	return count
}
