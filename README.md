<h1 align = "center">prim 即时通讯框架</h1>

### 简介

prim即时通讯框架是基于gowebsocket框架用go语言写的一个支持无限纵向扩展，企业级的即使通讯框架

### 架构设计

1. prim从设计上讲只是一个***中立***的，它内部不会涉及到任何的业务代码，它的架构是纯粹的。  
2. prim框架内部通过mongodb维护一个***企业账号***机制，每一个企业账号会构建一个沙盒空间，里面所有的websocket的coon连接是相通的，不同的企业账号具有不同的空间。所以prim从设计思路上讲它更像一个平台，帮不同的企业做即时通讯。  

### prim 架构图

![prim架构图](https://srillia.oss-cn-beijing.aliyuncs.com/即时通讯系统架构.jpg?versionId=CAEQDRiBgMCx1OruwxciIDRkMjI1YjczZGRkZTRkNWI4NmM4Mjg3OTVjZGZjMTg3)

### 架构图解

1. 从**prim架构图**上，可以看出来，prim相对比较中立，只维护基本的websocket连接,具体的业务逻辑是有**第三方服务器**提供。

2. **第三方服务器**获取临时key的过程，就是**第三方服务器**使用在**prim**注册的企业账号获取一个**prim**认证的token,然后**业务系统**把token给前端，前端使用token与prim建立一个连接。
3. prim自身的架构里，使用**mongodb**作为它自己的存储，包括，企业账号的逻辑。

### prim 框架接口

**系统用户**就是上面说的**企业用户**

```
func InitWebRouters() {
	router = gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 用户组
	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", user.List)
		userRouter.POST("/online", user.Online) //TODO 需要改请求
		userRouter.GET("/listUserMaps", user.ListUserMap)
		//userRouter.POST("/sendMessage", user.SendMessage)
		//userRouter.POST("/sendMessageAll", user.SendMessageAll)
		//userRouter.GET("/sendMessageTest", user.SendMessageTest)
	}

	// 消息
	msgRouter := router.Group("/msg")
	{
		msgRouter.POST("/sendMsg", msg.SendMsg) //TODO 需要改请求
		//获取消息的方法 @author Fran
		msgRouter.GET("/getMsg", msg.GetMsg)
	}

	//系统账户信息
	sysClientRouter := router.Group("/sysClient")
	{
		sysClientRouter.POST("/createSysClient", sysClient.CreateSysClient)
		sysClientRouter.GET("/getSysClient", sysClient.GetSysClient)
		sysClientRouter.POST("/getToken", sysClient.GetToken)
		//加上用户退出，使用退出websocket连接；
		sysClientRouter.POST("/exit", sysClient.ExitClient)
		sysClientRouter.POST("/updateTokenVal", sysClient.UpdateTokenVal)
	}

	// 系统
	systemRouter := router.Group("/system")
	{
		systemRouter.GET("/state", systems.Status)
	}

}

```


**Websocket 路由**

```
func WebsocketInit() {
	websocket.Register("heartbeat", websocket.HeartbeatController)
	websocket.Register("exit", websocket.ExitController)
	websocket.Register("ping", websocket.PingController)
	websocket.Register("msg", websocket.MsgController)
}
```

**websocket连接** 

```
ws://192.168.10.235:8089/acc
```

### prim框架实体

**企业用户**

```
type SysClient struct {
	PhoneNum string `json:"phoneNum.omitempty"`
	Account  string `json:"account.omitempty"`
	Password string `json:"password.omitempty"`
}
```

**用户**

```
type User struct {
	AccIp         string `json:"accIp"`         // acc Ip
	AccPort       string `json:"accPort"`       // acc 端口
	AppPlatform   string `json:"appPlatform"`   // appPlatform
	SysAccount    string `json:"sysAccount"`    // 系统账户
	UserId        string `json:"userId"`        // 用户Id
	Avatar        string `json:"avatar"`        // avatar 头像
	Nickname      string `json:"nickname"`      // nickname 昵称
	ClientIp      string `json:"clientIp"`      // 客户端Ip
	ClientPort    string `json:"clientPort"`    // 客户端端口
	LoginTime     uint64 `json:"loginTime"`     // 用户上次登录时间
	HeartbeatTime uint64 `json:"heartbeatTime"` // 用户上次心跳时间
	LogOutTime    uint64 `json:"logOutTime"`    // 用户退出登录的时间
	Address       string `json:"address"`       // 用户地址
	Qua           string `json:"qua"`           // qua
	DeviceInfo    string `json:"deviceInfo"`    // 设备信息
	IsLogoff      bool   `json:"isLogoff"`      // 是否下线
}

```

### main.go (Prim框架所有的逻辑都在main.go里面)

```
/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 09:59
 */

package main

import (
	"prim/initialize"
	"prim/routers"
	"prim/servers/grpcserver"
	"prim/servers/task"
	"prim/servers/websocket"
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

```
