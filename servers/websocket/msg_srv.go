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
	"prim/initialize"
	"prim/lib/cache"
	"prim/models"
	"prim/servers/grpcclient"
	"time"
)

//给所有的receiverIds发送信息，如果receiverIds为一个，为一对一聊天
func SendMessageToReceivers(client *Client, acc *models.Acc, receiverIds []string) (sendResults bool, err error) {
	sendResults = true
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("所有服务器正常", err)

		return
	}

	for _, server := range servers {
		if initialize.IsLocal(server) {
			SendMessagesLocally(client, acc, receiverIds)
		} else {
			accJson, err := json.Marshal(acc)
			if err != nil {
				//todo 处理异常
			}
			grpcclient.SendMsg(server, client.SysAccount, client.AppPlatform, receiverIds, accJson)
		}
	}

	return
}

//向所有的receiverIds发送信息
func SendMessagesLocally(client *Client, acc *models.Acc, receiverIds []string) {

	//获取redis 保存的userKey的格式
	userKeys := GetUserKeysNeedMsging(client.SysAccount, client.AppPlatform, client.UserId, receiverIds)
	for _, userKey := range userKeys {
		if client, ok := clientManager.Users[userKey]; ok {
			message, _ := json.Marshal(acc)
			client.SendMsg(message)
		} else {
			// todo: 如果对方不在线，将离线数据存到mongo中,临时保存
		}
	}
}
