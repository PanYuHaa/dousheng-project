package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// 每次启动服务器的时候都要对其进行初始化，用token来找userId，从而保证数据的安全性（暂时将token永久保存在sql中）
var usersLoginInfo = map[string]int64{}

var userIdSequence int64

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
			Id:       userIdSequence,
			NickName: username,
		}
		newAccount := model.Account{
			Id:       userIdSequence,
			UserName: username,
			PassWord: password,
			Token:    token,
		}
		usersLoginInfo[token] = newUser.Id             // 更新map
		err := service.AddAccount(newAccount, newUser) // 注册账号，更新sql数据
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

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	//t := service.IsAccountExist(username)
	//if t {
	//	result := service.ComparePassword(password, username)
	//	if !result {
	//		c.JSON(http.StatusOK, model.UserLoginResponse{
	//			Response: model.Response{StatusCode: 1, StatusMsg: "Password Wrong"},
	//		})
	//	} else {
	//		token := username + password
	//		c.JSON(http.StatusOK, model.UserLoginResponse{
	//			Response: model.Response{StatusCode: 0, StatusMsg: "Login Success"},
	//			UserId:   service.GetUserIdByName(username),
	//			Token:    token,
	//		})
	//
	//	}
	if id, exist := usersLoginInfo[token]; exist {
		if service.InfoVerify(password, username, id) {
			c.JSON(http.StatusOK, model.UserLoginResponse{
				Response: model.Response{StatusCode: 0, StatusMsg: "Login success"},
				UserId:   id,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, model.UserLoginResponse{
				Response: model.Response{StatusCode: 1, StatusMsg: "Incorrect username or password"},
			})
		}
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
