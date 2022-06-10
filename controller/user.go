package controller

import (
	"dousheng-demo/middleware"
	"dousheng-demo/model"
	"dousheng-demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sync/atomic"
)

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

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

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不一样，只需保存一份便可

	if _, exist := usersLoginInfo[username]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		// 创建账户，生成用户
		atomic.AddInt64(&userIdSequence, 1)
		newAccount := model.Account{
			Id:       userIdSequence,
			UserName: username,
			PassWord: encodePWD,
		}
		newUser := model.User{
			UserId: userIdSequence,
			Name:   username,
		}
		usersLoginInfo[username] = newUser          // 将生成的用户装进map，username作为key
		err := service.AddUser(newUser, newAccount) // 拉取当前账户及其用户的信息，并存储到数据库
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Failed to save"},
			})
			return
		}
		// 签发token
		newClaim := middleware.UserClaims{
			Id:   userIdSequence,
			Name: username,
		}
		token := middleware.GenerateToken(&newClaim)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Success"},
			UserId:   userIdSequence,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// 校验账户密码（保证账户的唯一性）
	if !service.IdentityVerify(username, password) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Username or password failed"},
		})
		return
	}

	// map缓存所有用户对象，并用username作为key，通过唯一的username来查找User对象
	if user, exist := usersLoginInfo[username]; exist {
		// 签发token，目前没有加入过期校验
		newClaim := middleware.UserClaims{
			Id:   user.UserId,
			Name: username,
		}
		token := middleware.GenerateToken(&newClaim)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Success"},
			UserId:   user.UserId,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	// token := c.Query("token")
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)

	if user, exist := usersLoginInfo[claim.Name]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Success"},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
