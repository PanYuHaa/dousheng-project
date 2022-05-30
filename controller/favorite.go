package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	//UserId := c.Query("user_id")
	VideoId := c.Query("video_id")
	token := c.Query("token")
	ActionType := c.Query("action_type")
	if t, exist := usersLoginInfo[token]; exist {
		UserId := strconv.FormatInt(t.Id, 10)
		if ActionType == "1" {
			err := service.Favorite(UserId, VideoId)
			if err != nil {
				c.JSON(http.StatusOK, model.UserLoginResponse{
					Response: model.Response{StatusCode: 1, StatusMsg: "Favorite failed"},
				})
			}
			c.JSON(http.StatusOK, model.UserLoginResponse{
				Response: model.Response{StatusCode: 0, StatusMsg: "Favorite success"},
			})
		} else if ActionType == "2" {
			err := service.Disfavorite(UserId, VideoId)
			if err != nil {
				c.JSON(http.StatusOK, model.UserLoginResponse{
					Response: model.Response{StatusCode: 1, StatusMsg: "Disfavorite failed"},
				})
			}
			c.JSON(http.StatusOK, model.UserLoginResponse{
				Response: model.Response{StatusCode: 0, StatusMsg: "Disfavorite success"},
			})
		}
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User don't exist"},
		})
	}
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	UserId := c.Query("user_id")
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, FeedResponse{
			VideoList: service.FavoriteList(UserId),
			Response:  service.FavoriteListRsp(UserId),
		})
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User don't login"},
		})
	}
}
