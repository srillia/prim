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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"prim/common"
	"prim/initialize"
	"prim/lib/cache"
	"prim/lib/mongolib"
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
			SendMessagesByLocal(client, acc, receiverIds)
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
func SendMessagesByLocal(client *Client, acc *models.Acc, receiverIds []string) {

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

func SaveMessageInMongo(client *Client, msg *models.Msg) {
	checkRoomInfoMongo(client, msg)
	addMessageInMongo(client, msg)
}

func disposeMsgAcc(client *Client, acc *models.Acc) (msg *models.Msg) {
	msg = &models.Msg{}
	common.ConvertType(acc.Msg, msg)
	msg.DateTime = common.TimestampToDateString(msg.Time)
	userOncline, _ := cache.GetUserOnlineInfo(client.GetKey())
	//判断空值，如果，用户
	if msg.SenderInfo.UserId == "" {
		msg.SenderInfo.UserId = userOncline.UserId
	}
	if msg.SenderInfo.Avatar == "" {
		msg.SenderInfo.Avatar = userOncline.Avatar
	}
	if msg.SenderInfo.Nickname == "" {
		msg.SenderInfo.Nickname = userOncline.Nickname
	}
	acc.Msg = msg
	return
}

func addMessageInMongo(client *Client, msg *models.Msg) {
	primMessage := &models.PrimMessage{}
	common.CopyProperties(msg, primMessage)
	primMessage.SysAccount = client.SysAccount
	mongolib.InsertOne(mongolib.GetConn("prim_message"), primMessage)
}

func checkRoomInfoMongo(client *Client, msgData *models.Msg) {
	if msgData.RoomId == "" {
		addRoomInfoWhenNoRoomIdInMsg(client, msgData)

	} else {
		objectId, _ := primitive.ObjectIDFromHex(msgData.RoomId)
		primRoom := &models.PrimRoom{}
		err := mongolib.FindOne(mongolib.GetConn("prim_room"), bson.M{"_id": objectId}, primRoom)
		//说明不存在
		if err != nil {
			addRoomInfoWhenNoRoomIdInMsg(client, msgData)
		}
		//存在不不用处理了
	}
}

func addRoomInfoWhenNoRoomIdInMsg(client *Client, msgData *models.Msg) {
	userList := [2]string{client.UserId, msgData.ReceiverId}
	primRoom := &models.PrimRoom{}
	err := mongolib.FindOne(mongolib.GetConn("prim_room"), bson.D{{"userList", bson.D{{"$all", userList}}}}, primRoom)
	//err == nil说明存在
	if err == nil {
		msgData.RoomId = primRoom.Id.Hex()
	} else { //不存在room 添加并赋值
		room := models.PrimRoom{UserList: userList}
		id, _ := mongolib.InsertOne(mongolib.GetConn("prim_room"), room)
		msgData.RoomId = id.(primitive.ObjectID).Hex()
	}
}
