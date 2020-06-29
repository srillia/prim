/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-30
* Time: 12:27
 */

package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"prim/lib/cache"
	"prim/models"
	"prim/servers/grpcclient"
	"time"
)

// 查询所有用户
func UserList() (userList []string) {

	userList = make([]string, 0)
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		var (
			list []string
		)
		if IsLocal(server) {
			list = GetUserList()
		} else {
			list, _ = grpcclient.GetUserList(server)
		}
		userList = append(userList, list...)
	}

	return
}

// 查询所有用户
func UserListByServer() (userMapByServer map[string][]string) {

	userMapByServer = make(map[string][]string)
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		userMapByServer[server.Ip] = make([]string, 0, 4)
		if IsLocal(server) {
			userMapByServer[server.Ip] = append(userMapByServer[server.Ip], GetUserList()...)
		} else {
			list, _ := grpcclient.GetUserList(server)
			userMapByServer[server.Ip] = append(userMapByServer[server.Ip], list...)
		}
	}
	return
}

// 查询用户是否在线
func CheckUserOnline(sysAccount string, appPlatform string, userId string) (online bool) {
	// 全平台查询
	if appPlatform == "all" {
		for _, platform := range GetAppPlatforms() {
			online, _ = checkUserOnline(sysAccount, platform, userId)
			if online == true {
				break
			}
		}
	} else {
		online, _ = checkUserOnline(sysAccount, appPlatform, userId)
	}

	return
}

// 查询用户 是否在线
func checkUserOnline(sysAccount string, appPlatform string, userId string) (online bool, err error) {
	key := GetUserKey(sysAccount, appPlatform, userId)
	userOnline, err := cache.GetUserOnlineInfo(key)
	if err != nil {
		if err == redis.Nil {
			fmt.Println("GetUserOnlineInfo", sysAccount, userId, err)

			return false, nil
		}

		fmt.Println("GetUserOnlineInfo", sysAccount, userId, err)

		return
	}

	online = userOnline.IsOnline()

	return
}

//给所有的receiverIds发送信息，如果receiverIds为一个，为一对一聊天
func SendMessageToReceivers(seq string, message interface{}, client *Client, receiverIds []string) (sendResults bool, err error) {
	sendResults = true
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("所有服务器正常", err)

		return
	}

	for _, server := range servers {
		if IsLocal(server) {
			SendMessagesLocally(message, client, receiverIds)
		} else {
			msgJson, err := json.Marshal(message)
			if err != nil {
				//todo 处理异常
			}
			grpcclient.SendMsg(server, seq, client.SysAccount, client.AppPlatform, receiverIds, msgJson)
		}
	}

	return
}

//向所有的receiverIds发送信息
func SendMessagesLocally(msgData interface{}, client *Client, receiverIds []string) {

	//获取redis 保存的userKey的格式
	userKeys := GetUserKeysNeedMsging(client.SysAccount, client.AppPlatform, client.UserId, receiverIds)
	for _, userKey := range userKeys {
		if client, ok := clientManager.Users[userKey]; ok {
			message, _ := json.Marshal(msgData)
			client.SendMsg(message)
		} else {
			// todo: 如果对方不在线，将离线数据存到mongo中,临时保存
		}
	}
}

// 给用户发送消息
//func SendUserMessage(sysAccount string, appPlatform string, userId string, msgId, message string) (sendResults bool, err error) {
//
//	data := models.GetTextMsgData(userId, msgId, message)
//
//	// TODO::需要判断不在本机的情况
//	sendResults, err = SendUserMessageLocal(sysAccount, appPlatform, userId, data)
//	if err != nil {
//		fmt.Println("给用户发送消息", appPlatform, userId, err)
//	}
//
//	return
//}

// 给本机用户发送消息
func SendUserMessageLocal(sysAccount string, appPlatform string, userIds []string, data []byte) {

	msg := &models.Msg{}
	json.Unmarshal(data, msg)

	for _, userId := range userIds {

		client := GetUserClient(sysAccount, appPlatform, userId)
		if client != nil {
			fmt.Println("给对应的用户发送信息")
			// 发送消息
			client.SendMsg(data)
		}
	}
}

// 给全体用户发消息
func SendUserMessageAll(sysAccount string, appPlatform string, userId string, msgId, cmd, message string) (sendResults bool, err error) {
	sendResults = true

	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		if IsLocal(server) {
			data := models.GetMsgData(userId, msgId, cmd, message)
			AllSendMessages(sysAccount, appPlatform, userId, data)
		} else {
			grpcclient.SendMsgAll(server, msgId, appPlatform, userId, cmd, message)
		}
	}

	return
}
