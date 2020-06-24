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
	"prim/helper"
	"prim/lib/cache"
	"prim/lib/redislib"
	"prim/models"
	"time"
)

var (
	clientManager = NewClientManager()                // 管理者
	appPlatforms  = []string{"web", "android", "ios"} // 全部的平台

	serverIp   string
	serverPort string
)

func GetAppPlatforms() []string {
	return appPlatforms
}

func GetClientManager() *ClientManager {
	return clientManager
}

func GetServer() (server *models.Server) {
	server = models.NewServer(serverIp, serverPort)

	return
}

func IsLocal(server *models.Server) (isLocal bool) {
	if server.Ip == serverIp && server.Port == serverPort {
		isLocal = true
	}

	return
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

	serverIp = helper.GetServerIp()

	webSocketPort := viper.GetString("app.webSocketPort")
	//rpcPort := viper.GetString("app.rpcPort")

	serverPort = webSocketPort

	http.HandleFunc("/acc", wsPage)

	// 添加处理程序
	go clientManager.start()
	fmt.Println("WebSocket 启动程序成功", serverIp, serverPort)

	http.ListenAndServe(":"+webSocketPort, nil)
}

func wsPage(w http.ResponseWriter, req *http.Request) {
	token := req.URL.Query().Get("token")

	var arr []string
	var pass bool
	//todo 现在的加密不是特别安全
	if pass, arr = redislib.PassCheckToken(token); pass == false {
		//直接return,没有权限连接
		return
	}
	sysAccount := arr[0]
	appPlatform := arr[1]
	userId := arr[2]

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

func saveUserAndClient(client *Client, currentTime uint64) {

	if !InAppPlatforms(client.AppPlatform) {
		fmt.Println("用户登录 不支持的平台", client.AppPlatform)
		//todo 不支持客户端逻辑处理
		return
	}

	// 存储用户登录数据
	userOnline := models.UserLogin(serverIp, serverPort, client.AppPlatform, client.SysAccount, client.UserId, client.Addr, currentTime)
	err := cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		fmt.Printf("保存登录用户信息成功:%v", userOnline)

	}

	// 连接存在，在添加,在缓存添加用户client信息
	if clientManager.InClient(client) {
		userKey := GetUserKey(client.SysAccount, client.AppPlatform, client.UserId)
		clientManager.AddUsers(userKey, client)
	}

}
