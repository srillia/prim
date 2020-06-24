/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-01
* Time: 10:40
 */

//todo 这个类将作废
package models

import (
	"prim/common"
)

const (
	MessageTypeText = "text"

	MessageCmdMsg   = "msg"
	MessageCmdEnter = "enter"
	MessageCmdExit  = "exit"
)

// 消息的定义
type Message struct {
	Target string `json:"target"` // 目标
	Type   string `json:"type"`   // 消息类型 text/img/
	Msg    string `json:"msg"`    // 消息内容
	From   string `json:"from"`   // 发送者
}

////todo 进一步规范消息格式
//type Msg struct {
//	SysAccount  string
//	SenderId    string
//	ReceiverId  string
//	time        time.Time
//	Message     string
//	RoomId      string
//	MsyType     string //1信息，2系统消息，3机器人，4
//	MsyContType string //1,文本，2 emoj,3,文件，4图片，5语音，6视频
//}

func NewTestMsg(from string, Msg string) (message *Message) {

	message = &Message{
		Type: MessageTypeText,
		From: from,
		Msg:  Msg,
	}

	return
}

func getTextMsgData(cmd, uuId, msgId, message string) string {
	textMsg := NewTestMsg(uuId, message)
	head := NewResponseHead(msgId, cmd, common.OK, "Ok", textMsg)

	return head.String()
}

// 文本消息
func GetMsgData(uuId, msgId, cmd, message string) string {

	return getTextMsgData(cmd, uuId, msgId, message)
}

// 文本消息
func GetTextMsgData(uuId, msgId, message string) string {

	return getTextMsgData("msg", uuId, msgId, message)
}

// 用户进入消息
func GetTextMsgDataEnter(uuId, msgId, message string) string {

	return getTextMsgData("enter", uuId, msgId, message)
}

// 用户退出消息
func GetTextMsgDataExit(uuId, msgId, message string) string {

	return getTextMsgData("exit", uuId, msgId, message)
}
