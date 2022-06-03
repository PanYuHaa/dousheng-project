package controller

import (
	"dousheng-demo/middleware"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	//UserId := c.Query("user_id")
	VideoId := c.Query("video_id")
	//token := c.Query("token")
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)
	ActionType := c.Query("action_type")
	if t, exist := usersLoginInfo[claim.Name]; exist {
		UserId := strconv.FormatInt(t.Id, 10)
		if ActionType == "1" {
			if !service.SearchFavorite(UserId, VideoId) {
				err := service.Favorite(UserId, VideoId)
				if err != nil {
					c.JSON(http.StatusOK, UserLoginResponse{
						Response: Response{StatusCode: 1, StatusMsg: "Favorite failed"},
					})
				}
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 0, StatusMsg: "Favorite success"},
				})
			} else {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "Favorite failed"},
				})
			}
		} else if ActionType == "2" {
			if service.SearchFavorite(UserId, VideoId) {
				err := service.Disfavorite(UserId, VideoId)
				if err != nil {
					c.JSON(http.StatusOK, UserLoginResponse{
						Response: Response{StatusCode: 1, StatusMsg: "Disfavorite failed"},
					})
				}
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 0, StatusMsg: "Disfavorite success"},
				})
			} else {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "Disfavorite failed"},
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User don't exist"},
		})
	}
}

func FavoriteList(c *gin.Context) {
	//token := c.Query("token")
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)
	UserId := c.Query("user_id")
	if _, exist := usersLoginInfo[claim.Name]; exist {
		if service.FavoriteListRsp() {
			c.JSON(http.StatusOK, FeedResponse{
				VideoList: service.FavoriteList(UserId),
				Response:  Response{StatusCode: 1, StatusMsg: "Favorite video"},
			})
			return
		}
		c.JSON(http.StatusOK, FeedResponse{
			VideoList: service.FavoriteList(UserId),
			Response:  Response{StatusCode: -1, StatusMsg: "No Favorite Videos"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User don't login"},
		})
	}
}
