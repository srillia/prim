/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:20
 */

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"prim/controllers/home"
	"prim/controllers/sysclient"
	"prim/controllers/systems"
	"prim/controllers/user"
	"prim/helper"
)

var (
	router *gin.Engine
)

func InitWebRouters() {
	router = gin.Default()
	router.LoadHTMLGlob("views/**/*")

	//router.LoadHTMLFiles("views/**/*")

	// 用户组
	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", user.List)
		userRouter.GET("/online", user.Online)
		userRouter.GET("/listUserMaps", user.ListUserMap)
		//userRouter.POST("/sendMessage", user.SendMessage)
		//userRouter.POST("/sendMessageAll", user.SendMessageAll)
		//userRouter.GET("/sendMessageTest", user.SendMessageTest)
	}

	// 用户组
	sysClientRouter := router.Group("/sysClient")
	{
		sysClientRouter.POST("/createSysClient", sysClient.CreateSysClient)
		sysClientRouter.GET("/getSysClient", sysClient.GetSysClient)
		sysClientRouter.GET("/getToken", sysClient.GetToken)
	}

	// 系统
	systemRouter := router.Group("/system")
	{
		systemRouter.GET("/state", systems.Status)
	}

	// home
	homeRouter := router.Group("/home")
	{
		homeRouter.GET("/index", home.Index)
	}

	// router.POST("/user/online", user.Online)
}

func InitHttpServer() {
	httpPort := viper.GetString("app.httpPort")
	fmt.Println("Http Server 启动成功", helper.GetServerIp(), httpPort)
	http.ListenAndServe(":"+httpPort, router)

}
