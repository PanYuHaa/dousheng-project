package controller

import (
	"dousheng-demo/middleware"
	"dousheng-demo/model"
	"dousheng-demo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowListResponse struct {
	Response
	FollowList []model.User `json:"video_list,omitempty"`
	NextTime   int64        `json:"next_time,omitempty"`
}

func RelationAction(c *gin.Context) {
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)
	toid := c.Query("to_user_id")
	ActionType := c.Query("action_type")
	if t, exist := usersLoginInfo[claim.Name]; exist {
		t := strconv.FormatInt(t.UserId, 10)
		if ActionType == "1" {
			if !service.SearchFollow(t, toid) {
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
			} else {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "Already Follow"},
				})
			}
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

func FollowList(c *gin.Context) {
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)
	userId := c.Query("user_id")
	if _, exist := usersLoginInfo[claim.Name]; exist {
		if service.FollowListRsp() {
			c.JSON(http.StatusOK, FollowListResponse{
				FollowList: service.FollowList(userId),
				Response:   Response{StatusCode: 1, StatusMsg: "FollowList"},
			})
			return
		}
		c.JSON(http.StatusOK, FollowListResponse{
			FollowList: service.FollowList(userId),
			Response:   Response{StatusCode: -1, StatusMsg: "No FollowList"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User don't login"},
		})
	}
}
