package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
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
