package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func AddComment(comment model.Comment, videoId int64) error {
	err := repository.AddComment(comment, videoId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentId int64, videoId int64) error {
	err := repository.DeleteCommentById(commentId, videoId)
	if err != nil {
		return err
	}
	return nil
}
