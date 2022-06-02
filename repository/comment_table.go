package repository

import (
	"dousheng-demo/model"
	"gorm.io/gorm"
)

//	Add

func AddComment(comment model.Comment, videoId int64) error {
	commentId := comment.Id
	// 开启事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&comment).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		} // 创建评论

		newCommentVideo := model.VideoComment{
			VideoId:   videoId,
			CommentId: commentId,
		}
		if err := tx.Create(&newCommentVideo).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		} // 创建评论视频id关系表

		if err := tx.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

//	Delete

func DeleteCommentById(commentId int64, videoId int64) error {
	// 开启事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Where("id = ?", commentId).Delete(&model.Comment{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		} // 删除评论

		if err := tx.Where("comment_id = ?", commentId).Where("video_id = ?", videoId).Delete(&model.VideoComment{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		} // 删除评论视频关系表

		if err := tx.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

//	Get

func GetCommentCount(videoId int64) int64 {
	var video model.Video
	if err := DB.Where("id = ?", videoId).Find(&video).Error; err != nil {
		panic(err)
	}
	return video.CommentCount
}

func GetCommentIds(videoId int64) []int64 {
	var commentIds []int64
	if err := DB.Raw("select comment_id from video_comments where video_id=? order by comment_id desc", videoId).Scan(&commentIds).Error; err != nil {
		panic(err)
	}
	return commentIds
}

func GetCommentById(commentId int64) model.Comment {
	var comment model.Comment
	//if err := DB.Model(&model.Comment{}).Find(&comment, commentId); err != nil {
	//	panic(err)
	//}
	DB.Model(&model.Comment{}).Find(&comment, commentId)
	return comment
}
