/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 14:41
 */

package models

import "time"

/************************  请求数据  **************************/
// 通用请求数据格式
type Request struct {
	Seq  string      `json:"seq"`            // 消息的唯一Id
	Cmd  string      `json:"cmd"`            // 请求命令字
	Data interface{} `json:"data,omitempty"` // 数据 json
}

// 登录请求数据
type Login struct {
	ServiceToken string `json:"serviceToken"` // 验证用户是否登录
	AppId        uint32 `json:"appId,omitempty"`
	UserId       string `json:"userId,omitempty"`
}

// 登录请求数据
type Msg struct {
	AppId       uint32   `json:"appId"` // 验证用户是否登录
	SenderId    string   `json:"senderId,omitempty"`
	ReceiverIds []string `json:"receiverIds,omitempty"`
	time        uint64   `json:"time,omitempty"`
	Message     string   `json:"message,omitempty"`
	Type        string   `json:"type,omitempty"` // img ,text
}

type SingleMsg struct {
	SysAccount  string
	SenderId    string
	ReceiverId  string
	time        time.Time
	Message     string
	RoomId      string
	MsyType     byte //1信息，2系统消息，3机器人，4
	MsyContType byte //1,文本，2 emoj,3,文件，4图片，5语音，6视频
}

type GroupMsg struct {
	//todo
}

// 心跳请求数据
type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}
