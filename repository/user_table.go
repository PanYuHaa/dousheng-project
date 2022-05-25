package repository

import (
	"dousheng-demo/model"
)

var DemoUser = model.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

func GetUserById(userId int64) model.User {
	// 从db中获取user
	var user model.User
	DB.Model(&model.User{}).Find(&user, userId)
	return user
}
