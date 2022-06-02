package model

type Favorite struct {
	UserId  string
	VideoId string
}

type Follow struct {
	UserId   string
	ToUserId string
}

type VideoComment struct {
	VideoId   int64
	CommentId int64
}
