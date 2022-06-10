package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
	"os"
	"os/exec"
	"sync/atomic"
	"time"
)

var videoId = int64(4)

func PublishVideo(user model.User, finalName string, title string) error {
	// playUrl := "http://[$主机ip]:8080/static/video/" + finalName
	playUrl := "http://124.223.184.55:8080/static/video/" + finalName
	err := getSnapshot(finalName)
	if err != nil {
		return err
	}
	// coverUrl := "http://[$主机ip]:8080/static/video/" + finalName
	coverUrl := "http://192.168.10.103:8080/static/cover/" + getPicName(finalName)
	atomic.AddInt64(&videoId, 1) // 原子增加防止出现脏读
	return repository.AddVideo(model.Video{
		//Id: repository.VideoAmount() + 1,
		Id: videoId,
		Author: model.User{
			UserId: user.UserId,
			Name:   user.Name,
		},
		PlayUrl:    playUrl,
		CoverUrl:   coverUrl,
		Title:      title,
		CreateTime: time.Now().Unix(),
	})
}

func getSnapshot(finalName string) error {
	outPicName := getPicName(finalName)
	inputPath := "./public/video/" + finalName
	outputPath := "./public/cover/" + outPicName
	// 调用ffmpeg应用程序进行视频截图
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ss", "1", "-f", "image2", "-frames:v", "1", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 由于有些ffmpeg的bug没有解决所以暂时不反回error
	return cmd.Run()
}

func getPicName(finalName string) string {
	var filename string
	for i, val := range finalName {
		if val == '.' {
			filename = finalName[:i]
			break
		}
	}
	return filename + ".png" // 照片名
}
