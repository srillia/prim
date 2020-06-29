/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:24
 */

package websocket

import (
	"fmt"
	"prim/helper"
	"prim/lib/cache"
	"prim/models"
	"sync"
	"time"
)

// 连接管理
type ClientManager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户 // appPlatform+uuid
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}

	return
}

// 获取用户key
func GetUserKey(account string, appPlatform string, userId string) (key string) {
	key = fmt.Sprintf("%s_%s_%s", account, appPlatform, userId)

	return
}

// 获取用户key
func GetUserKeysInAllPlatform(account string, userId string) (keys []string) {
	for _, platform := range GetAppPlatforms() {
		keys = append(keys, fmt.Sprintf("%s_%s_%s", account, platform, userId))
	}
	return
}

// 获取用户key
func GetUserKeysExcludeThePlatform(account string, appPlatform string, userId string) (keys []string) {
	for _, platform := range GetAppPlatforms() {
		if platform != appPlatform {
			keys = append(keys, fmt.Sprintf("%s_%s_%s", account, platform, userId))
		}
	}
	return
}

// 获取用户key
func GetUserKeysNeedMsging(account string, appPlatform string, senderId string, receiverIds []string) (keys []string) {
	//先获取我其它平台的用户，这些用户要需要接受信息
	keys = GetUserKeysExcludeThePlatform(account, appPlatform, senderId)
	//添加所有需要发送的用记的信息，并添加到keys 切片返回
	for _, receiverId := range receiverIds {
		for _, receiverIdInPlatform := range GetUserKeysInAllPlatform(account, receiverId) {
			keys = append(keys, receiverIdInPlatform)
		}

	}
	return
}

/**************************  manager  ***************************************/

func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	// 连接存在，在添加
	_, ok = manager.Clients[client]

	return
}

// GetClients
func (manager *ClientManager) GetClients() (clients map[*Client]bool) {

	clients = make(map[*Client]bool)

	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value

		return true
	})

	return
}

// 遍历
func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {

	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	for key, value := range manager.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}

	return
}

// GetClientsLen
func (manager *ClientManager) GetClientsLen() (clientsLen int) {

	clientsLen = len(manager.Clients)

	return
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	manager.Clients[client] = true
}

// 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	if _, ok := manager.Clients[client]; ok {
		delete(manager.Clients, client)
	}
}

// 获取用户的连接
func (manager *ClientManager) GetUserClient(sysAccount string, appPlatform string, userId string) (client *Client) {

	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()

	userKey := GetUserKey(sysAccount, appPlatform, userId)
	if value, ok := manager.Users[userKey]; ok {
		client = value
	}

	return
}

// GetClientsLen
func (manager *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(manager.Users)

	return
}

// 添加用户
func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()

	manager.Users[key] = client
}

// 删除用户
func (manager *ClientManager) DelUsers(client *Client) (result bool) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()

	key := GetUserKey(client.SysAccount, client.AppPlatform, client.UserId)
	if value, ok := manager.Users[key]; ok {
		// 判断是否为相同的用户
		if value.Addr != client.Addr {

			return
		}
		delete(manager.Users, key)
		result = true
	}

	return
}

// 获取用户的key
func (manager *ClientManager) GetUserKeys() (userKeys []string) {

	userKeys = make([]string, 0)
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	for key := range manager.Users {
		userKeys = append(userKeys, key)
	}

	return
}

// 获取用户列表
func (manager *ClientManager) GetUserList() (userList []string) {

	userList = make([]string, 0)

	clientManager.UserLock.RLock()
	defer clientManager.UserLock.RUnlock()

	for _, v := range clientManager.Users {
		userList = append(userList, v.UserId)
		fmt.Println("GetUserList", v.AppPlatform, v.UserId, v.Addr)
	}

	fmt.Println("GetUserList", clientManager.Users)

	return
}

// 获取用户的客户端
func (manager *ClientManager) GetUserClients() (clients []*Client) {

	clients = make([]*Client, 0)
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	for _, v := range manager.Users {
		clients = append(clients, v)
	}

	return
}

// 向全部成员(除了自己)发送数据
func (manager *ClientManager) sendAll(message []byte, ignore *Client) {

	clients := manager.GetUserClients()
	for _, conn := range clients {
		if conn != ignore {
			conn.SendMsg(message)
		}
	}
}

// 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)

	fmt.Println("EventRegister 用户建立连接", client.Addr)

	// client.Send <- []byte("连接成功")
}

//todo 作废代码
//func (manager *ClientManager) EventLogin(login *login) {
//
//	client := login.Client
//	// 连接存在，在添加
//	if manager.InClient(client) {
//		userKey := login.GetKey()
//		manager.AddUsers(userKey, login.Client)
//	}
//
//	fmt.Println("EventLogin 用户登录", client.Addr, login.AppPlatform, login.UserId)
//
//	orderId := helper.GetOrderIdTime()
//	SendUserMessageAll(login.AppPlatform, login.UserId, orderId, models.MessageActionEnter, "哈喽~")
//}

// 用户断开连接
func (manager *ClientManager) EventUnregister(client *Client) {
	manager.DelClients(client)

	// 删除用户连接
	deleteResult := manager.DelUsers(client)
	if deleteResult == false {
		// 不是当前连接的客户端

		return
	}

	// 清除redis登录数据
	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err == nil {
		userOnline.LogOut()
		cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	}

	// 关闭 chan
	// close(client.Send)

	fmt.Println("EventUnregister 用户断开连接", client.Addr, client.AppPlatform, client.UserId)

	if client.UserId != "" {
		orderId := helper.GetOrderIdTime()
		SendUserMessageAll(client.SysAccount, client.AppPlatform, client.UserId, orderId, models.MessageActionExit, "用户已经离开~")
	}
}

// 管道处理程序
func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)

		//todo 下面代码作废
		//case login := <-manager.Login:
		//	// 用户登录
		//	manager.EventLogin(login)

		case conn := <-manager.Unregister:
			//todo 用户离开通知处理
			println(conn)
			// 断开连接事件
			//manager.EventUnregister(conn)

		case message := <-manager.Broadcast:
			// 广播事件
			clients := manager.GetClients()
			for conn := range clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
				}
			}
		}
	}
}

/**************************  manager info  ***************************************/
// 获取管理者信息
func GetManagerInfo(isDebug string) (managerInfo map[string]interface{}) {
	managerInfo = make(map[string]interface{})

	managerInfo["clientsLen"] = clientManager.GetClientsLen()
	managerInfo["usersLen"] = clientManager.GetUsersLen()
	managerInfo["chanRegisterLen"] = len(clientManager.Register)
	//managerInfo["chanLoginLen"] = len(clientManager.Login)
	managerInfo["chanUnregisterLen"] = len(clientManager.Unregister)
	managerInfo["chanBroadcastLen"] = len(clientManager.Broadcast)

	if isDebug == "true" {
		addrList := make([]string, 0)
		clientManager.ClientsRange(func(client *Client, value bool) (result bool) {
			addrList = append(addrList, client.Addr)

			return true
		})

		users := clientManager.GetUserKeys()

		managerInfo["clients"] = addrList
		managerInfo["users"] = users
	}

	return
}

//TODO 获取在线用户的信息
// 获取用户所在的连接
func GetUserClient(sysAccount string, appPlatform string, userId string) (client *Client) {
	client = clientManager.GetUserClient(sysAccount, appPlatform, userId)

	return
}

// 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())

	clients := clientManager.GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			fmt.Println("心跳时间超时 关闭连接", client.Addr, client.UserId, client.LoginTime, client.HeartbeatTime)

			client.Socket.Close()
		}
	}
}

// 获取全部用户
func GetUserList() (userList []string) {
	fmt.Println("获取全部用户")

	userList = clientManager.GetUserList()

	return
}

// 全员广播
func AllSendMessages(sysAccount string, appPlatform string, userId string, data string) {
	fmt.Println("全员广播", sysAccount, appPlatform, userId, data)

	ignore := clientManager.GetUserClient(sysAccount, appPlatform, userId)
	clientManager.sendAll([]byte(data), ignore)
}
