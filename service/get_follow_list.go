package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetFollowList(userId int64) []model.User {
	var id int64
	followList := make([]model.User, 0)
	followIds := repository.GetUserFollow(userId)
	for _, id = range followIds {
		followList = append(followList, repository.GetUserById(id))
	}
	return followList
}
