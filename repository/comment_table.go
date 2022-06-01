package repository

import (
	"dousheng-demo/model"
	"gorm.io/gorm"
)

// Add

func AddComment(comment model.Comment, videoId int64) error {
	// 开启事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&comment).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

// Delete

func DeleteCommentById(commentId int64, videoId int64) error {
	// 开启事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Where("id = ?", commentId).Delete(&model.Comment{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}
