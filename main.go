package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple-douyin/config"
	"simple-douyin/model"
	"strconv"
)

func main() {
	// 下面这一行是聊天服务、需要写再启动、并且把services文件夹拷贝过来
	//go service.RunMessageServer()

	_, err := model.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	initRouter(r)

	configLocal := config.GetConfig()
	host := configLocal.ServerHost
	port := configLocal.ServerPort

	err = r.Run(host + ":" + strconv.Itoa(port))
	if err != nil {
		return
	}
}
