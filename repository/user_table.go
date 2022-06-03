package repository

import (
	"dousheng-demo/model"
	"gorm.io/gorm"
)

//	Add

func AddUser(user model.User, account model.Account) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Model(&model.User{}).Create(&user).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Model(&model.Account{}).Create(&account).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

//	Get

func GetUserById(userId int64) model.User {
	// 从db中获取user
	var user model.User
	DB.Table("users").Find(&user, userId)
	return user
}

func GetAccount(username string, password string) model.Account {
	var account model.Account
	DB.Table("accounts").Where("user_name = ?", username).Where("pass_word = ?", password).Find(&account)
	return account
}

func GetUsersAmount() int64 {
	// 从db中获取user的数量(ID)
	var count int64
	DB.Model(&model.User{}).Where("nick_name != ?", "").Count(&count)
	return count
}
