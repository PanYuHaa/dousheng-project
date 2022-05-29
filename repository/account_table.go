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
