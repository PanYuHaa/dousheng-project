package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

// var startId = repository.TimeLimitAmount(9999999999) // 初始化startId
var startId int64
var count int64

func GetVideoList() []model.Video {
	if startId == 0 {
		return nil
	}

	videoList := make([]model.Video, 0)
	count = 0 // 初始化计数器
	for {
		videoList = append(videoList, repository.GetVideoById(startId))
		count++
		startId--
		// 如果list满30个，返回videoList
		if count == 30 {
			startId = repository.TimeLimitAmount(videoList[29].CreateTime)
			return videoList
		}
		// 如果原数据库中已经没有video，返回videoList
		if startId == 0 {
			startId = repository.TimeLimitAmount(9999999999)
			return videoList
		}
	}
}

func GetVideoRsp() model.Response {
	if startId == 0 {
		return model.Response{StatusCode: -1, StatusMsg: "No video"}
	} else {
		return model.Response{StatusCode: 0, StatusMsg: "Success"}
	}
}

func GetCreateTime() int64 {
	if GetVideoList() == nil {
		return 0
	} else {
		lastId := count - 1
		return GetVideoList()[lastId].CreateTime // 获取最后一个播放视频的创建时间，用作下次提取list的开始
	}
}
