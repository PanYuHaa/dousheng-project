package model

type User struct {
	Id            int64  `json:"id,omitempty"`
	NickName      string `json:"nick_name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Account struct {
	Id       int64  `json:"id,omitempty"`
	UserName string `gorm:"column:username" json:"username,omitempty"`
	PassWord string `gorm:"column:password" json:"password,omitempty"`
}
