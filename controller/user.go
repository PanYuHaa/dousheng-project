package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
var usersLoginInfo = map[string]model.User{}

// test data: username=ceshi, password=douyin
//var usersLoginInfo = map[string]model.User{
//	"ceshidouyin": {
//		Id:            1,
//		Name:          "ceshi",
//		FollowCount:   10,
//		FollowerCount: 5,
//		IsFollow:      true,
//	},
//}

var userIdSequence = int64(1)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := model.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0, StatusMsg: "Success"},
			UserId:   userIdSequence,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password // 使用加密方法生成token

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0, StatusMsg: "Success"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		err := service.AddUser(user) // 拉取当前登录用户的全部信息，并存储到本地
		if err != nil {
			c.JSON(http.StatusOK, model.UserResponse{
				Response: model.Response{StatusCode: 1, StatusMsg: "Failed to save"},
			})
			return
		}
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 0, StatusMsg: "Success"},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
