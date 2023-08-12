package main

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/model"
)

func main() {
	// 下面这一行是聊天服务、需要写再启动、并且把services文件夹拷贝过来
	//go service.RunMessageServer()
	err := model.InitDB()
	if err != nil {
		return
	}

	r := gin.Default()
	initRouter(r)

	err = r.Run()
	if err != nil {
		return
	}
}
