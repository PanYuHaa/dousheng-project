package repository

import (
	"dousheng-demo/model"
)

func GetUserById(userId int64) model.User {
	// 从db中获取user
	var user model.User
	DB.Table("user_infos").Find(&user, userId)
	return user
}

func IsUserExist(username string) bool {
	var user model.User
	DB.Table("user_infos").Where("name = ?", username).Find(&user)
	if user.NickName == username {
		return true
	}
	return false
}

func GetUsersAmount() int64 {
	// 从db中获取user的数量(ID)
	var count int64
	DB.Model(&model.Video{}).Where("name = ?", "").Count(&count)
	return count
}

func AddAccount(account model.Account) error {
	dbRes := DB.Model(&model.Account{}).Create(&account)
	return dbRes.Error
}
