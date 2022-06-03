package controller

import (
	"dousheng-demo/middleware"
	"dousheng-demo/model"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment model.Comment `json:"comment,omitempty"`
}

var commentIdSequence = int64(0)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	//token := c.Query("token")
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[claim.Name]; exist {
		if actionType == "1" {
			atomic.AddInt64(&commentIdSequence, 1)
			text := c.Query("comment_text")
			videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
			t := time.Now()
			newComment := model.Comment{
				Id:         commentIdSequence,
				User:       user,
				Content:    text,
				CreateDate: strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day()),
			}
			service.AddComment(newComment, videoId)
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Comment success"},
				Comment:  newComment,
			})
			return
		}
		if actionType == "2" {
			commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
			service.DeleteComment(commentId, videoId)
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Delete comment success"},
			})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Failed to comment action"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	//token := c.Query("token")
	userClaim, _ := c.Get("userClaim")
	claim := userClaim.(*middleware.UserClaims)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if _, exist := usersLoginInfo[claim.Name]; exist {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: service.GetCommentList(videoId),
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
