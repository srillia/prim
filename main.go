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
	"io"
	"net/http"
	"os"
	"os/exec"
	"prim/lib/redislib"
	"prim/routers"
	"prim/servers/grpcserver"
	"prim/servers/task"
	"prim/servers/websocket"
	"time"
)

func main() {
	initConfig()

	initFile()

	initRedis()

	router := gin.Default()
	// 初始化路由
	routers.Init(router)
	routers.WebsocketInit()

	// 定时任务
	task.Init()

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

// 初始化日志
func initFile() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}

func initConfig() {

	runEnv := os.Getenv("RUN_ENV")

	switch runEnv {
	case "local":
		viper.SetConfigName("config/local")
	case "dev":
		viper.SetConfigName("config/dev")
	case "test":
		viper.SetConfigName("config/test")
	case "gray":
		viper.SetConfigName("config/gray")
	case "prod":
		viper.SetConfigName("config/prod")
	default:
		viper.SetConfigName("config/local")
	}
	viper.AddConfigPath(".") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config redis:", viper.Get("redis"))

}

func initRedis() {
	redislib.ExampleNewClient()
}

func open() {

	time.Sleep(1000 * time.Millisecond)

	httpUrl := viper.GetString("app.httpUrl")
	httpUrl = "http://" + httpUrl + "/home/index"

	fmt.Println("访问页面体验:", httpUrl)

	cmd := exec.Command("open", httpUrl)
	cmd.Output()
}
