package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	FollwList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	if !service.GetVideoRsp() {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: -1, StatusMsg: "No video"},
			FollwList: service.GetVideoList(),
			NextTime:  service.GetCreateTime(),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "Success"},
		FollwList: service.GetVideoList(),
		NextTime:  service.GetCreateTime(),
	})
}
