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
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"prim/common"
	"prim/lib/cache"
	"prim/lib/mongolib"
	"prim/models"
	"time"
)

// ping
func PingController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	fmt.Println("webSocket_request ping接口", client.Addr, seq, message)

	data = "pong"

	return
}

// 给用户发送消息
func MsgController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {
	code = common.OK

	msgData := &models.Msg{}
	if err := json.Unmarshal(message, msgData); err != nil {
		code = common.ParameterIllegal
		fmt.Println("用户登录 解析数据失败", seq, err)

		return
	}
	msgData.SysAccount = client.SysAccount
	msgData.SenderId = client.UserId
	msgData.DateTime = common.TimestampToDateString(msgData.Time)
	checkRoomInfoMongo(client, msgData)
	addMessageInMongo(msgData)

	fmt.Printf("给用户发送消息:%v", msgData)

	SendMessageToReceivers(msgData, client, []string{msgData.ReceiverId})
	return
}

func addMessageInMongo(msgData *models.Msg) {
	mongolib.InsertOne(mongolib.GetConn("prim_message"), msgData)
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
		msgData.RoomId = id.(primitive.ObjectID).String()
	}
}

//todo 此方法已经做废
//func LoginController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {
//	code = common.OK
//	currentTime := uint64(time.Now().Unix())
//
//	request := &models.Login{}
//	if err := json.Unmarshal(message, request); err != nil {
//		code = common.ParameterIllegal
//		fmt.Println("用户登录 解析数据失败", seq, err)
//
//		return
//	}
//
//	fmt.Println("webSocket_request 用户登录", seq, "ServiceToken", request.ServiceToken)
//	// TODO 用户登录需要区分是否是客服。
//	// TODO::进行用户权限认证，一般是客户端传入TOKEN，然后检验TOKEN是否合法，通过TOKEN解析出来用户ID
//	// 本项目只是演示，所以直接过去客户端传入的用户ID
//	if request.UserId == "" || len(request.UserId) >= 20 {
//		code = common.UnauthorizedUserId
//		fmt.Println("用户登录 非法的用户", seq, request.UserId)
//
//		return
//	}
//
//	if !InAppPlatforms(request.AppPlatform) {
//		code = common.Unauthorized
//		fmt.Println("用户登录 不支持的平台", seq, request.AppPlatform)
//
//		return
//	}
//
//	if client.IsLogin() {
//		fmt.Println("用户登录 用户已经登录", client.AppPlatform, client.UserId, seq)
//		code = common.OperationFailure
//
//		return
//	}
//
//	client.Login(request.AppPlatform, request.UserId, currentTime)
//
//	// 存储数据
//	userOnline := models.UserLogin(serverIp, serverPort, request.AppPlatform, request.UserId, client.Addr, currentTime)
//	err := cache.SetUserOnlineInfo(client.GetKey(), userOnline)
//	if err != nil {
//		code = common.ServerError
//		fmt.Println("用户登录 SetUserOnlineInfo", seq, err)
//
//		return
//	}
//
//	// 用户登录
//	login := &login{
//		AppPlatform:  request.AppPlatform,
//		UserId: request.UserId,
//		Client: client,
//	}
//	clientManager.Login <- login
//
//	fmt.Println("用户登录 成功", seq, client.Addr, request.UserId)
//
//	return
//}

// 心跳接口
func HeartbeatController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := uint64(time.Now().Unix())

	request := &models.HeartBeat{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		fmt.Println("心跳接口 解析数据失败", seq, err)

		return
	}

	fmt.Println("webSocket_request 心跳接口", client.AppPlatform, client.UserId)

	if !client.IsLogin() {
		fmt.Println("心跳接口 用户未登录", client.AppPlatform, client.UserId, seq)
		code = common.NotLoggedIn

		return
	}

	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err != nil {
		if err == redis.Nil {
			code = common.NotLoggedIn
			fmt.Println("心跳接口 用户未登录", seq, client.AppPlatform, client.UserId)

			return
		} else {
			code = common.ServerError
			fmt.Println("心跳接口 GetUserOnlineInfo", seq, client.AppPlatform, client.UserId, err)

			return
		}
	}

	client.Heartbeat(currentTime)
	userOnline.Heartbeat(currentTime)
	err = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		fmt.Println("心跳接口 SetUserOnlineInfo", seq, client.AppPlatform, client.UserId, err)

		return
	}

	return
}
