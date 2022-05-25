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
	userIdStr := c.Query("user_id")
	var userId int64
	userId, _ = strconv.ParseInt(userIdStr, 10, 64)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: service.GetPublishList(userId + 1), // 因为下标号和数据库的id不对应
	})
}
