package controller

import (
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toid := c.Query("to_user_id")
	ActionType := c.Query("action_type")
	if t, exist := usersLoginInfo[token]; exist {
		t := strconv.FormatInt(t.Id, 10)
		if ActionType == "1" {
			err := service.Follow(t, toid)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "Follow failed"},
				})
				return
			}
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Follow success"},
			})
		} else if ActionType == "2" {
			err := service.UnFollow(t, toid)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "Unfollow failed"},
				})
				return
			}
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Unfollow success"},
			})
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User don't exist"},
		})
	}
}
