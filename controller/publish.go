package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/video", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// 用user_id来找
	//userIdStr := c.Query("user_id")
	//var userId int64
	//userId, _ = strconv.ParseInt(userIdStr, 10, 64)

	// 用token来找，因为登录状态下才有PublishList
	token := c.Query("token")
	userId := usersLoginInfo[token].Id

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: service.GetPublishList(userId),
	})
}
