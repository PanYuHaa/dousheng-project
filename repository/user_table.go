package repository

import (
	"dousheng-demo/model"
)

//	Add

func AddUser(user model.User) error {
	dbRes := DB.Model(&model.User{}).Create(&user)
	return dbRes.Error
}

//	Get

func GetUserById(userId int64) model.User {
	// 从db中获取user
	var user model.User
	DB.Table("users").Find(&user, userId)
	return user
}

func GetUsersAmount() int64 {
	// 从db中获取user的数量(ID)
	var count int64
	DB.Model(&model.User{}).Where("nick_name != ?", "").Count(&count)
	return count
}
