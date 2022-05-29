package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetLastUserId() int64 {
	return repository.GetUsersAmount()
}

func AddAccount(account model.Account, user model.User) error {
	err1 := repository.AddAccount(account)
	if err1 != nil {
		return err1
	}
	err2 := repository.AddUser(user)
	if err2 != nil {
		// 如果增加账号成功了但是增加初始化后的用户信息失败，要回滚一下账户的数据库
		repository.DeleteAccountById(account.Id)
		return err2
	}
	return nil
}
