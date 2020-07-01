/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:11
 */

package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"prim/common"
	"prim/controllers"
	"prim/servers/websocket"
	"strconv"
)

// 查看全部在线用户
func List(c *gin.Context) {

	//todo 此处需要鉴权

	data := make(map[string]interface{})

	userList := websocket.UserList()
	data["userList"] = userList

	controllers.Response(c, common.OK, "", data)
}

// 查看全部在线用户
func ListUserMap(c *gin.Context) {

	//todo 此处需要鉴权

	data := make(map[string]interface{})

	userMap := websocket.UserListByServer()
	data["userList"] = userMap

	controllers.Response(c, common.OK, "", data)
}

// 查看用户是否在线
func Online(c *gin.Context) {

	userId := c.Query("userId")
	appPlatformStr := c.Query("appPlatform")
	sysAccount := c.Query("sysAccount")

	fmt.Println("http_request 查看用户是否在线", userId, appPlatformStr)
	appPlatform, _ := strconv.ParseInt(appPlatformStr, 10, 32)

	data := make(map[string]interface{})

	online := websocket.CheckUserOnline(sysAccount, string(appPlatform), userId)
	data["userId"] = userId
	data["online"] = online

	controllers.Response(c, common.OK, "", data)
}

//// 给用户发送消息
//func SendMessage(c *gin.Context) {
//	// 获取参数
//	appPlatformStr := c.PostForm("appPlatform")
//	userId := c.PostForm("userId")
//	msgId := c.PostForm("msgId")
//	message := c.PostForm("message")
//
//	//拿到请求中的token
//	token := c.Request.Header.Get("authorization")
//	fmt.Print(token)
//	fmt.Println("http_request 给用户发送消息", appPlatformStr, userId, msgId, message)
//	appPlatform, _ := strconv.ParseInt(appPlatformStr, 10, 32)
//
//	// TODO::进行用户权限认证，一般是客户端传入TOKEN，然后检验TOKEN是否合法，通过TOKEN解析出来用户ID
//	// 本项目只是演示，所以直接过去客户端传入的用户ID(userId)
//
//	data := make(map[string]interface{})
//
//	if cache.SeqDuplicates(msgId) {
//		fmt.Println("给用户发送消息 重复提交:", msgId)
//		controllers.Result(c, common.OK, "", data)
//
//		return
//	}
//
//	sendResults, err := websocket.SendUserMessage(string(appPlatform), userId, msgId, message)
//	if err != nil {
//		data["sendResultsErr"] = err.Error()
//	}
//
//	data["sendResults"] = sendResults
//
//	controllers.Result(c, common.OK, "", data)
//}
//
//// 给全员发送消息
//func SendMessageAll(c *gin.Context) {
//	// 获取参数
//	appPlatformStr := c.PostForm("appPlatform")
//	userId := c.PostForm("userId")
//	msgId := c.PostForm("msgId")
//	message := c.PostForm("message")
//
//	fmt.Println("http_request 给全体用户发送消息", appPlatformStr, userId, msgId, message)
//
//	appPlatform, _ := strconv.ParseInt(appPlatformStr, 10, 32)
//
//	data := make(map[string]interface{})
//	if cache.SeqDuplicates(msgId) {
//		fmt.Println("给用户发送消息 重复提交:", msgId)
//		controllers.Result(c, common.OK, "", data)
//
//		return
//	}
//
//	sendResults, err := websocket.SendUserMessageAll(string(appPlatform), userId, msgId, models.MessageActionMsg, message)
//	if err != nil {
//		data["sendResultsErr"] = err.Error()
//
//	}
//
//	data["sendResults"] = sendResults
//
//	controllers.Result(c, common.OK, "", data)
//
//}
//
//func SendMessageTest(c *gin.Context) {
//	fmt.Println("进入测试方法", c)
//	token := c.Request.Header.Get("authorization")
//	fmt.Println("获取的token为：", token)
//	val, flag := c.GetPostForm("token")
//	if flag {
//		fmt.Println("获取的值:", val)
//	}
//
//	controllers.Result(c, common.OK, val, nil)
//}
