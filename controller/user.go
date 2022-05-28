package controller

import (
	"dousheng-demo/model"
	"dousheng-demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func Login(c *gin.Context) {
	fmt.Printf("get here")
	username := c.Query("username")
	password := c.Query("password")
	t := service.IsAccountExist(username)
	if t {
		result := service.ComparePassword(password, username)
		if !result {
			c.JSON(http.StatusOK, model.UserLoginResponse{
				Response: model.Response{StatusCode: 1, StatusMsg: "Password Wrong"},
			})
		} else {
			token := username + password
			c.JSON(http.StatusOK, model.UserLoginResponse{
				Response: model.Response{StatusCode: 0, StatusMsg: "Login Success"},
				UserId:   service.GetUserIdByName(username),
				Token:    token,
			})
		}
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	var userid int64
	userid, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)
	fmt.Println(userid)
	if service.IsAccountExistById(userid) {
		c.JSON(http.StatusOK, model.UserInfo{
			Response:      model.Response{StatusCode: 0, StatusMsg: "Success"},
			FollowCount:   service.GetUserFollowCountByID(userid),
			FollowerCount: service.GetUserFollowerCountByID(userid),
			ID:            userid,
			IsFollow:      service.GetUserIsFollowByID(userid),
			Name:          service.GetUserNameByID(userid),
		})
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
