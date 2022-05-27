package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

var userIdSequence int64

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if service.IsAccountExist(username) {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1) // 对addr指向的值加上delta，再返回*addr
		newAccount := model.Account{
			Id:       userIdSequence,
			UserName: username,
			PassWord: password,
		}
		newUser := model.User{
			Id:       userIdSequence,
			NickName: username,
		}
		err := service.AddAccount(newAccount, newUser) // 注册账号
		if err != nil {
			c.JSON(http.StatusOK, model.UserLoginResponse{
				Response: model.Response{StatusCode: -1, StatusMsg: "Register Account failed"},
			})
			return
		}
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0, StatusMsg: "Success"},
			UserId:   userIdSequence,
			Token:    token,
		})
	}
}

//func UserInfo(c *gin.Context) {
//	token := c.Query("token")
//
//	if user, exist := usersLoginInfo[token]; exist {
//		c.JSON(http.StatusOK, model.UserResponse{
//			Response: model.Response{StatusCode: 0},
//			User:     user,
//		})
//	} else {
//		c.JSON(http.StatusOK, model.UserResponse{
//			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
//		})
//	}
//}