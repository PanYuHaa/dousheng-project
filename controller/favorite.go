package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(c *gin.Context) {
	UserId := c.Query("user_id")
	VideoId := c.Query("video_id")
	token := c.Query("token")
	ActionType := c.Query("action_type")
	if _, exist := usersLoginInfo[token]; exist {
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
