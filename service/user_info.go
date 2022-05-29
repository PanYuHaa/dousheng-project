package service

import (
	"dousheng-demo/repository"
)

func IsTokenMatch(userid int64, token string) bool {
	return repository.IsTokenMatch(userid, token)
}

func IsUserExistById(id int64) bool {
	return repository.IsUserExistById(id)
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
