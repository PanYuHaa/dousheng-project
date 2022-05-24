package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
	"time"
)

func GetVideoList() []model.Video{
	if repository.GetVideoAmount() == 0 {
		return nil
	}
	// 单次提取视频有30个限制
	var len int64
	if repository.GetVideoAmount() > 30{
		len = 30
	} else {
		len = repository.GetVideoAmount()
	}
	videoList := make([]model.Video, len)
	j := 0
	for i := repository.GetVideoAmount(); i > 0; i--{
		videoList[j] = repository.GetVideo(i)
		j++
	}
	return videoList
}

func GetVideoRsp() model.Response{
	if repository.GetVideoAmount() == 0 {
		return model.Response{StatusCode: -1}
	}else {
		return model.Response{StatusCode: 0, StatusMsg: "Success."}
	}
}

func GetVideoTime() int64 {
	if GetVideoList() == nil {
		return 0
	} else {
		return time.Now().Unix()
	}
}