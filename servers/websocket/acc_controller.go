/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 13:12
 */

package websocket

import (
	"fmt"
	"prim/lib/cache"
	"prim/models"
	"time"
)

// ping
func ExitController(client *Client, acc *models.Acc) *models.Acc {
	ClearClient(client, acc)
	return acc.ExitAcc(nil)
}

// ping
func PingController(client *Client, acc *models.Acc) *models.Acc {
	fmt.Println("webSocket_request ping接口", client.Addr, acc.Seq, acc.Msg)
	return acc.OkAcc(models.Ok{Result: "pong"})
}

// 给用户发送消息
func MsgController(client *Client, acc *models.Acc) *models.Acc {
	msg := disposeMsgAcc(client, acc)
	SaveMessageInMongo(client, msg)
	fmt.Printf("给用户发送消息:%v", msg)
	SendMessageToReceivers(client, acc, []string{msg.ReceiverId})
	return acc.AckAcc(nil)
}

// 心跳接口
func HeartbeatController(client *Client, acc *models.Acc) *models.Acc {

	currentTime := uint64(time.Now().Unix())

	fmt.Println("webSocket_request 心跳接口", client.AppPlatform, client.UserId)

	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err != nil {
		fmt.Println("心跳接口 用户不在线", acc.Seq, client.SysAccount, client.AppPlatform, client.UserId)
	}

	client.Heartbeat(currentTime)
	userOnline.Heartbeat(currentTime)
	err = cache.SetUserOnlineInfo(client.GetKey(), userOnline)

	return acc.HeartBeatAcc(models.HeartBeat{Result: "pulse success"})
}
