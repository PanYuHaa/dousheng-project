package main

import (
	"dousheng-demo/controller"
	"dousheng-demo/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin") // 路由组加入中间件

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)                                      // PYH
	apiRouter.GET("/user/", middleware.JwtVerify, controller.UserInfo)            // CHY
	apiRouter.POST("/user/register/", controller.Register)                        // PYH
	apiRouter.POST("/user/login/", controller.Login)                              // CMD
	apiRouter.POST("/publish/action/", controller.Publish)                        // PYH
	apiRouter.GET("/publish/list/", middleware.JwtVerify, controller.PublishList) // PYH

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JwtVerify, controller.FavoriteAction) //CMD
	apiRouter.GET("/favorite/list/", middleware.JwtVerify, controller.FavoriteList)      //CMD
	apiRouter.POST("/comment/action/", middleware.JwtVerify, controller.CommentAction)   //PYH
	apiRouter.GET("/comment/list/", middleware.JwtVerify, controller.CommentList)        //PYH

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.JwtVerify, controller.RelationAction) //CHY
	apiRouter.GET("/relation/follow/list/", middleware.JwtVerify, controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
