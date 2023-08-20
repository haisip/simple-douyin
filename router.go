package main

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/controller"
	"simple-douyin/middleware"
)

func initRouter(r *gin.Engine) {
	// 公共文件夹文件目录
	// todo 配置文件添加 public文件夹路径
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// todo 不需要认证的API添加到 apiRouter
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	// todo token 作为参数但是不是非必须参数的及接口
	aApiRouter := apiRouter.Group("/", middleware.TokenAuthMiddleware(false))
	aApiRouter.GET("/feed/", controller.Feed)

	// todo 使用中间件做用户认证，需要token的API请添加到protectedApiRouter中
	protectedApiRouter := apiRouter.Group("/", middleware.TokenAuthMiddleware(true))
	protectedApiRouter.GET("/user/", controller.UserInfo)
	protectedApiRouter.GET("/publish/list/", controller.PublishList)
	protectedApiRouter.POST("/publish/action/", controller.Publish)

	// 添加 handler
	//protectedApiRouter.POST("/favorite/list/", controller.FavoriteList)

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
