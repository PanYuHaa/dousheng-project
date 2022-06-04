package repository

import "dousheng-demo/model"

func AddNewFollow(subscribe model.Follow) error {
	mu.Lock()
	defer mu.Unlock()
	dbRes := DB.Model(&model.Follow{}).Create(&subscribe)
	var user, touser model.User
	DB.Table("users").Find(&user, subscribe.UserId)
	user.FollowCount++
	DB.Save(&user)
	DB.Table("users").Find(&touser, subscribe.ToUserId)
	touser.FollowerCount++
	DB.Save(&touser)
	return dbRes.Error
}

func DeleteFollow(subscribe model.Follow) error {
	mu.Lock()
	defer mu.Unlock()
	dbRes := DB.Where("to_user_id = ? ", subscribe.ToUserId).Where("user_id = ?", subscribe.UserId).Delete(&model.Follow{})
	var user, touser model.User
	DB.Table("users").Find(&user, subscribe.UserId)
	user.FollowCount--
	DB.Save(&user)
	DB.Table("users").Find(&touser, subscribe.ToUserId)
	touser.FollowerCount--
	DB.Save(&touser)
	return dbRes.Error
}

func SearchFollow(subscribe model.Follow) bool {
	var t model.Follow
	t.UserId = "-1"
	DB.Where("to_user_id = ? ", subscribe.ToUserId).Where("user_id = ?", subscribe.UserId).Find(&t)
	if t.UserId == "-1" {
		return false
	} else {
		return true
	}
}

func GetUserFollow(UserId string) []string {
	var Ids []string
	DB.Raw("select * from users where user_id=?", UserId).Scan(&Ids)
	return Ids
}
