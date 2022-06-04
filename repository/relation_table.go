package repository

import (
	"dousheng-demo/model"
	"gorm.io/gorm"
)

func AddNewFollow(subscribe model.Follow) error {
	mu.Lock()
	defer mu.Unlock()
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Follow{}).Create(&subscribe).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.User{}).Where("user_id = ?", subscribe.ToUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.User{}).Where("user_id = ?", subscribe.UserId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

func DeleteFollow(subscribe model.Follow) error {
	mu.Lock()
	defer mu.Unlock()
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Follow{}).Where("to_user_id = ? ", subscribe.ToUserId).Where("user_id = ?", subscribe.UserId).Delete(&subscribe).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.User{}).Where("user_id = ?", subscribe.ToUserId).Update("follower_count", gorm.Expr("follower_count + ?", -1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.User{}).Where("user_id = ?", subscribe.UserId).Update("follow_count", gorm.Expr("follow_count + ?", -1)).Error; err != nil {
			return err
		}
		return nil
	})
}

func SearchFollow(subscribe model.Follow) bool {
	var t model.Follow
	t.UserId = "-1"
	DB.Where("to_user_id = ? ", subscribe.ToUserId).Where("user_id = ?", subscribe.UserId).Find(&t)
	if t.UserId == "-1" {
		return false
	} else {
		return true
	}
}

func GetUserFollow(UserId string) []string {
	var Ids []string
	DB.Raw("select * from users where user_id=?", UserId).Scan(&Ids)
	return Ids
}
