package main

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/controller"
)

func initRouter(r *gin.Engine) {
	// todo 使用中间件做用户认证，需要token的API请添加到protectedApiRouter中

	// 公共文件夹文件目录
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	// todo 不需要认证的API添加到 apiRouter
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.GET("/user/", controller.UserInfo)
	//apiRouter.GET("/feed/", controller.Feed)

	// todo 需要认证的api 添加到 protectedApiRouter
	//protectedApiRouter := apiRouter.Group("/", middleware.TokenAuthMiddleware())
	// 添加 handler
	//protectedApiRouter.POST("/favorite/list/", controller.FavoriteList)
	//protectedApiRouter.POST("/publish/action/", controller.Publish)
	//protectedApiRouter.GET("/publish/list/", controller.PublishList)

	//// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	//apiRouter.GET("/relation/friend/list/", controller.FriendList)
	//apiRouter.GET("/message/chat/", controller.MessageChat)
	//apiRouter.POST("/message/action/", controller.MessageAction)
}
