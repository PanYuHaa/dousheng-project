package controller

import (
	"dousheng-demo/middleware"
	"dousheng-demo/model"
	"dousheng-demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	token := c.PostForm("token")
	if token == "" {
		panic("token not exist !")
	}
	claim := middleware.ParseToken(token)

	if _, exist := usersLoginInfo[claim.Name]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)                 // filepath.Base()返回路径最后一个元素
	user := usersLoginInfo[claim.Name]                       // token验证
	finalName := fmt.Sprintf("%d_%s", user.UserId, filename) // 格式化字符串拼接
	saveFile := filepath.Join("./public/video", finalName)   // filepath.Join()连接路径，saveFile文件保存的目标地址。
	// 存储视频在服务端
	if err = c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 更新video数据表
	if err = service.PublishVideo(user, finalName, title); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	//token := c.Query("token")
	// token校验
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)
	if _, exist := usersLoginInfo[claim.Name]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	userIdStr := c.Query("user_id")
	var userId int64
	userId, _ = strconv.ParseInt(userIdStr, 10, 64)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: service.GetPublishList(userId),
	})
}
