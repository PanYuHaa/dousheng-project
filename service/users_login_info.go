package service

import "dousheng-demo/repository"

func GetLoginInfo() map[string]int64 {
	var usersLoginInfo = map[string]int64{}
	// 导入token对应userid的map
	var id int64
	for id = 1; id <= repository.GetUsersAmount(); id++ {
		usersLoginInfo[repository.GetTokenById(id)] = id
	}
	return usersLoginInfo
}
