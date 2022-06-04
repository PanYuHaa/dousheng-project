package service

import (
	"dousheng-demo/model"
	"dousheng-demo/repository"
	"fmt"
	"strconv"
)

var i = 0
var userId string
var Id string

func FollowList(userId string) []model.User {
	FollowList := make([]model.User, 0)
	UserIds := repository.GetUserFollow(userId)
	for _, Id = range UserIds {
		i++
		Id64, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			fmt.Printf("wrong!")
		}
		FollowList = append(FollowList, repository.GetUserById(Id64))
	}

	return FollowList
}

func FollowListRsp() bool {
	if i != 0 {
		return true
	} else {
		return false
	}
}
