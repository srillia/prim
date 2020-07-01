/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:04
 */

package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
	"prim/initialize"
	"prim/lib/cache"
	"prim/lib/redislib"
	"prim/models"
	"prim/servers/grpcclient"
	"time"
)

var (
	clientManager = NewClientManager()                // 管理者
	appPlatforms  = []string{"web", "android", "ios"} // 全部的平台
)

func GetAppPlatforms() []string {
	return appPlatforms
}

func GetClientManager() *ClientManager {
	return clientManager
}

func InAppPlatforms(appPlatform string) (inAppPlatform bool) {

	for _, value := range appPlatforms {
		if value == appPlatform {
			inAppPlatform = true

			return
		}
	}

	return
}

// 启动程序
func StartWebSocket() {

	webSocketPort := viper.GetString("app.webSocketPort")

	http.HandleFunc("/acc", wsPage)

	// 添加处理程序
	go clientManager.start()
	fmt.Println("WebSocket 启动程序成功", initialize.ServerIp, webSocketPort)
	http.ListenAndServe(":"+webSocketPort, nil)
}

func wsPage(w http.ResponseWriter, req *http.Request) {
	token := req.URL.Query().Get("token")
	pass, sysAccount, appPlatform, userId := redislib.PassCheckToken(token)
	if pass == false {
		//直接return,没有权限连接
		return
	}
	//todo,先检测，clientManager里的client是否存在，存在删除，再建立连接

	clearExistsClient(sysAccount, appPlatform, userId)

	if !InAppPlatforms(appPlatform) {
		fmt.Println("用户登录 不支持的平台", appPlatform)
		//todo 不支持客户端逻辑处理
		return
	}

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, sysAccount, appPlatform, userId, currentTime)

	//注册用户客户端是可以用的
	clientManager.Register <- client

	//在redis保存
	saveUserAndClient(client, currentTime)

	go client.read()
	go client.write()

}

func clearExistsClient(sysAccount string, appPlatform string, userId string) {
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("所有服务器正常", err)

		return
	}

	for _, server := range servers {
		if initialize.IsLocal(server) {
			client := clientManager.GetUserClient(sysAccount, appPlatform, userId)
			ClearClient(client)
		} else {

			grpcclient.ClearExistsClient(server, sysAccount, appPlatform, userId)
		}
	}

}

func saveUserAndClient(client *Client, currentTime uint64) {

	// 存储用户登录数据
	userOnline := models.UserLogin(initialize.ServerIp, initialize.AccPort, client.AppPlatform, client.SysAccount, client.UserId, client.Addr, currentTime)
	err := cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		fmt.Printf("保存登录用户信息成功:%v", userOnline)

	}

	// 连接存在，在添加,在缓存添加用户client信息
	if clientManager.InClient(client) {
		userKey := GetUserKey(client.SysAccount, client.AppPlatform, client.UserId)
		clientManager.AddUser(userKey, client)
	}

}
