/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 09:59
 */

package main

import (
	"fmt"
	"github.com/spf13/viper"
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

	//初始化websocket路由
	routers.WebsocketInit()
	//初始化http路由，并启动http服务器
	routers.InitWebRouters()

	//go 关键字类似创建一个新的线程去执行其他内容，并行处理
	go websocket.StartWebSocket()
	// grpc
	go grpcserver.Init()
	// 定时任务
	go task.Init()

	//这个初始化之后，可以不调用
	//go open()

	//启动http服务器，阻塞线程
	routers.InitHttpServer()
}

func open() {

	time.Sleep(1000 * time.Millisecond)

	httpUrl := viper.GetString("app.httpUrl")
	httpUrl = "http://" + httpUrl + "/home/index"

	fmt.Println("访问页面体验:", httpUrl)

	action := exec.Command("open", httpUrl)
	action.Output()
}
