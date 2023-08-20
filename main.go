package main

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/config"
	"simple-douyin/model"
)

func main() {
	// 下面这一行是聊天服务、需要写再启动、并且把services文件夹拷贝过来
	//go service.RunMessageServer()

	// 下面这一行用于保存配置到json文件，但是似乎不用配置文件也行（使用config.go中的配置）
	_ = config.SaveConfigToFile()

	_, err := model.InitDB()
	if err != nil {
		return
	}

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	initRouter(r)

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}
