package controller

import (
	"dousheng-demo/middleware"
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		if !service.GetVideoRsp() {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 1, StatusMsg: "No video"},
			})
			return
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "Success"},
			VideoList: service.GetVideoList(),
			NextTime:  service.GetCreateTime(),
		})
		return
	}

	// token校验，用户状态
	claim := middleware.ParseToken(token)
	if _, exist := usersLoginInfo[claim.Name]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if !service.GetVideoRsp() {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: "No video"},
		})
		return
	}

	user := usersLoginInfo[claim.Name]
	userId := user.UserId
	videos := service.GetVideoList() // 视频列表
	// 处理显示是否关注
	followIds := service.GetUserFollow(userId) // 此用户id下关注的所有人
	for i := 0; i < len(videos); i++ {
		for j := 0; j < len(followIds); j++ {
			if videos[i].Author.UserId == followIds[j] {
				videos[i].Author.IsFollow = true
				break
			}
		}
	}
	// 处理显示是否喜欢
	favoriteIds := service.GetFavoriteVideos(userId)
	for i := 0; i < len(videos); i++ {
		for j := 0; j < len(favoriteIds); j++ {
			if videos[i].Id == favoriteIds[j] {
				videos[i].IsFavorite = true
				break
			}
		}
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "Success"},
		VideoList: videos,
		NextTime:  service.GetCreateTime(),
	})
}
