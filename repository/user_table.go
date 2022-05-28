package repository

import (
	"dousheng-demo/model"
)

func GetUserById(userId int64) model.User {
	// 从db中获取user
	var user model.User
	DB.Table("users").Find(&user, userId)
	return user
}

func IsAccountExist(username string) bool {
	var user model.Account
	DB.Table("accounts").Where("username = ?", username).Find(&user)
	if user.UserName == username {
		return true
	} else {
		return false
	}

}

func GetUsersAmount() int64 {
	// 从db中获取user的数量(ID)
	var count int64
	DB.Model(&model.User{}).Where("nick_name != ?", "").Count(&count)
	return count
}

func AddUser(user model.User) error {
	dbRes := DB.Model(&model.User{}).Create(&user)
	return dbRes.Error
}

func AddAccount(account model.Account) error {
	dbRes := DB.Model(&model.Account{}).Create(&account)
	return dbRes.Error
}

func DeleteAccountById(id int64) error {
	dbRes := DB.Delete(&model.Account{}, id)
	return dbRes.Error
}
func GetPasswordByUsername(username string) string {
	var account model.Account
	DB.Table("accounts").Where("username = ?", username).Find(&account)
	return account.PassWord
}
func GetUseridByName(username string) int64 {
	var account model.Account
	DB.Table("accounts").Where("username = ?", username).Find(&account)
	return account.Id
}
