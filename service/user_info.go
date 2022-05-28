package service

import (
	"dousheng-demo/repository"
)

func IsAccountExistById(userid int64) bool {
	return repository.IsAccountExistById(userid)
}

func GetUserFollowCountByID(userid int64) int64 {
	return repository.GetUserFollowCountByID(userid)
}

func GetUserFollowerCountByID(userid int64) int64 {
	return repository.GetUserFollowerCountByID(userid)
}

func GetUserIsFollowByID(userid int64) bool {
	return repository.GetUserIsFollowByID(userid)
}

func GetUserNameByID(userid int64) string {
	return repository.GetUserNameByID(userid)
}
