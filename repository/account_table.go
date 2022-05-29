package repository

import "dousheng-demo/model"

//	Add

func AddAccount(account model.Account) error {
	dbRes := DB.Model(&model.Account{}).Create(&account)
	return dbRes.Error
}

//	Delete

func DeleteAccountById(id int64) error {
	dbRes := DB.Delete(&model.Account{}, id)
	return dbRes.Error
}

//	Check

func IsTokenMatch(userid int64, token string) bool {
	var user model.Account
	DB.Table("accounts").Where("id = ?", userid).Find(&user)
	if user.Token == token {
		return true
	} else {
		return false
	}
}

func IsUserExistById(userid int64) bool {
	// user和account通过id查找的话建议都用此函数
	var user model.Account
	DB.Table("accounts").Where("id = ?", userid).Find(&user)
	if user.Id == userid {
		return true
	} else {
		return false
	}
}

//	Get

func GetPasswordById(id int64) string {
	var account model.Account
	DB.Table("accounts").Where("id = ?", id).Find(&account)
	return account.PassWord
}

func GetUsernameById(id int64) string {
	var account model.Account
	DB.Table("accounts").Where("id = ?", id).Find(&account)
	return account.UserName
}

func GetTokenById(id int64) string {
	var account model.Account
	DB.Table("accounts").Where("id = ?", id).Find(&account)
	return account.Token
}

func GetUserFollowCountByID(userid int64) int64 {
	var user model.User
	DB.Table("users").Find(&user, userid)
	return user.FollowCount
}

func GetUserFollowerCountByID(userid int64) int64 {
	var user model.User
	DB.Table("users").Find(&user, userid)
	return user.FollowerCount
}

func GetUserIsFollowByID(userid int64) bool {
	var user model.User
	DB.Table("users").Find(&user, userid)
	return user.IsFollow
}

func GetUserNameByID(userid int64) string {
	var user model.User
	DB.Table("users").Find(&user, userid)
	return user.NickName
}
