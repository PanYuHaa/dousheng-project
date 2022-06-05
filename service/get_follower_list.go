package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
)

func GetFollowerList(toUserId int64) []model.User {
	var id int64
	followerList := make([]model.User, 0)
	followerIds := repository.GetUserFollower(toUserId)
	for _, id = range followerIds {
		followerList = append(followerList, repository.GetUserById(id))
	}
	return followerList
}
