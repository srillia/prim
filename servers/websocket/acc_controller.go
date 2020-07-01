/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 13:12
 */

package websocket

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"prim/common"
	"prim/lib/cache"
	"prim/lib/mongolib"
	"prim/models"
	"time"
)

// ping
func ExitController(client *Client, acc *models.Acc) *models.Acc {
	ClearClient(client)
	return acc.OkAcc(nil)
}

// ping
func PingController(client *Client, acc *models.Acc) *models.Acc {
	fmt.Println("webSocket_request ping接口", client.Addr, acc.Seq, acc.Msg)
	return acc.OkAcc(models.Ok{Result: "pong"})
}

// 给用户发送消息
func MsgController(client *Client, acc *models.Acc) *models.Acc {

	msg := disposeMsgAcc(client, acc)
	checkRoomInfoMongo(client, msg)
	addMessageInMongo(client, msg)

	fmt.Printf("给用户发送消息:%v", msg)
	SendMessageToReceivers(client, acc, []string{msg.ReceiverId})
	return acc.AckAcc(nil)
}

func disposeMsgAcc(client *Client, acc *models.Acc) *models.Msg {
	msg := &models.Msg{}
	convertType(acc.Msg, msg)
	msg.SenderId = client.UserId
	msg.DateTime = common.TimestampToDateString(msg.Time)
	acc.Msg = msg
	return msg
}

func convertType(data interface{}, ret interface{}) {
	msgData, _ := json.Marshal(data)
	json.Unmarshal(msgData, ret)
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
		_, err := mongolib.FindOne(mongolib.GetConn("prim_room"), bson.M{"_id": objectId}, models.PrimRoom{})
		//说明不存在
		if err != nil {
			addRoomInfoWhenNoRoomIdInMsg(client, msgData)
		}
		//存在不不用处理了
	}
}

func addRoomInfoWhenNoRoomIdInMsg(client *Client, msgData *models.Msg) {
	userList := [2]string{client.UserId, msgData.ReceiverId}
	roomExist, err := mongolib.FindOne(mongolib.GetConn("prim_room"), bson.D{{"userlist", bson.D{{"$all", userList}}}}, models.PrimRoom{})
	//err == nil说明存在
	if err == nil {
		msgData.RoomId = roomExist.(models.PrimRoom).Id.String()
	} else { //不存在room 添加并赋值
		room := models.PrimRoom{UserList: userList}
		id, _ := mongolib.InsertOne(mongolib.GetConn("prim_room"), room)
		msgData.RoomId = id.(primitive.ObjectID).Hex()
	}
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
