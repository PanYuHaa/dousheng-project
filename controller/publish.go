package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// 用user_id来找
	userIdStr := c.Query("user_id")
	var userId int64
	userId, _ = strconv.ParseInt(userIdStr, 10, 64)
	// 用token来找
	//token := c.Query("token")
	//userId := usersLoginInfo[token].Id

	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: service.GetPublishList(userId),
	})
}
