package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetCommentList(videoId int64) []model.Comment {
	commentList := make([]model.Comment, 0)
	var i = int64(0)
	for {
		if i == repository.GetCommentCount(videoId) {
			return commentList
		}
		commentId := repository.GetCommentIds(videoId)[i]
		commentList = append(commentList, repository.GetCommentById(commentId))
		i++
	}
}
