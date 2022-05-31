package main

import (
	"dousheng-demo/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)                // PYH
	apiRouter.GET("/user/", controller.UserInfo)            // CHY
	apiRouter.POST("/user/register/", controller.Register)  // PYH
	apiRouter.POST("/user/login/", controller.Login)        // CMD
	apiRouter.POST("/publish/action/", controller.Publish)  // PYH
	apiRouter.GET("/publish/list/", controller.PublishList) // PYH

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction) //CMD
	apiRouter.GET("/favorite/list/", controller.FavoriteList)      //CMD
	//apiRouter.POST("/comment/action/", controller.CommentAction)	  //PYH
	//apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
