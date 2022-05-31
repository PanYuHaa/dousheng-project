package repository

import "dousheng-demo/model"

func AddNewFollow(subscribe model.Follow) error {
	dbRes := DB.Model(&model.Follow{}).Create(&subscribe)
	return dbRes.Error
}

func DeleteFollow(subscribe model.Follow) error {
	dbRes := DB.Where("to_user_id = ? ", subscribe.ToUserId).Where("user_id = ?", subscribe.UserId).Delete(&model.Follow{})
	return dbRes.Error
}
