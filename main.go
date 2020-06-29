/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 09:59
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/exec"
	"prim/initialize"
	"prim/routers"
	"prim/servers/grpcserver"
	"prim/servers/task"
	"prim/servers/websocket"
	"time"
)

func main() {
	initialize.InitConfig()

	initialize.InitFile()

	initialize.InitRedis()

	initialize.InitMongo()

	router := gin.Default()
	// 初始化路由
	routers.Init(router)
	routers.WebsocketInit()

	// 定时任务
	task.CleanInit()

	// 服务注册
	task.ServerInit()
	//go 关键字类似创建一个新的线程去执行其他内容，并行处理
	go websocket.StartWebSocket()
	// grpc
	go grpcserver.Init()
	//这个初始化之后，可以不调用
	go open()

	httpPort := viper.GetString("app.httpPort")
	http.ListenAndServe(":"+httpPort, router)

}

func open() {

	time.Sleep(1000 * time.Millisecond)

	httpUrl := viper.GetString("app.httpUrl")
	httpUrl = "http://" + httpUrl + "/home/index"

	fmt.Println("访问页面体验:", httpUrl)

	cmd := exec.Command("open", httpUrl)
	cmd.Output()
}
