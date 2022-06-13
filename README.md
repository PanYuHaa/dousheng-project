# dousheng-project

## 功能说明

接口已按照官方文档完善，接口如下:

基础接口:
1. 视频流接口
2. 用户注册
3. 用户登录
4. 用户信息
5. 投稿接口
6. 发布接口

扩展接口-I:\
7. 赞操作\
8. 点赞列表\
9. 评论操作\
10. 评论列表

扩展接口-II:\
11. 关注操作\
12. 关注列表\
13. 粉丝列表

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video/video_name 即可

## 代码结构
dousheng_project
├── config 
│   └── config.go // 设定基本配置信息（如数据库DSN）
├── controller // 控制层
│   ├── comment.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   ├── relation.go
│   ├── response.go
│   └── user.go
├── middleware // 中间件
│   └── jwt.go
├── model // 数据模型层
│   ├── comment.go
│   ├── realtion.go
│   ├── user.go
│   └── video.go
├── public // 视频与封面存储位置
│   ├── cover
│   └── video
├── repository // 数据持久层
│   ├── db_init.go
│   ├── comment_table.go
│   ├── favorite_table.go
│   ├── realtion_table.go
│   ├── user_table.go
│   └── video_table.go
├── service // 业务逻辑层
│   ├── comment_action.go
│   ├── favorite_action.go
│   ├── get_comment_list.go
│   ├── get_favorite_list.go
│   ├── get_follow_list.go
│   ├── get_follower_list.go
│   ├── get_publish_list.go
│   ├── get_video_list.go
│   ├── publish_video.go
│   ├── relation.go
│   ├── service_init.go
│   ├── user_login.go
│   └── user_register.go
├── go.mod
│   └── go.sum
├── main.go
└── router.go
