package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PrimSysClient struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNum string             `bson:"phoneNum,omitempty"`
	Account  string             `bson:"account,omitempty"`
	Password string             `bson:"password,omitempty"`
	State    byte               `bson:"state,omitempty"`
	AuthCode string             `bson:"authCode,omitempty"`
}

type PrimRoom struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserList [2]string          `bson:"userList,omitempty"`
}

type PrimMessage struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	SysAccount  string             `bson:"sysAccount,omitempty"`
	Time        int64              `json:"time,omitempty"`
	DateTime    string             `json:"dateTime,omitempty"`
	SenderInfo  SenderInfo         `bson:"senderInfo,omitempty"`
	ReceiverId  string             `bson:"receiverId,omitempty"`
	Message     string             `bson:"message,omitempty"`
	RoomId      string             `bson:"roomId,omitempty"`
	MsgType     string             `bson:"msgType,omitempty"`
	MsgContType string             `bson:"msgContType,omitempty"`
	State       string             `bson:"state,omitempty"`
}
