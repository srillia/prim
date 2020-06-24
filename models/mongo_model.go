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

//todo 下载注释代码将作废
//type PrimUser struct {
//	Id              primitive.ObjectID `bson:"_id,omitempty"`
//	UserId          string
//	PrimSysClientId primitive.ObjectID
//	Account         string
//	Nickname        string
//	Signature       string
//	Sex             int8
//	Birthday        time.Time
//	Tel             string
//	Email           string
//	Intro           string
//	Avatar          string
//	Age             int8
//	State           byte
//}

type PrimRoom struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserList [2]string          `bson:"userList,omitempty"`
}

type PrimMessage struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Message     string             `bson:"message,omitempty"`
	RoomId      primitive.ObjectID `bson:"roomId,omitempty"`
	SenderId    string             `bson:"roomId,omitempty"`
	ReceiverId  string             `bson:"receiverId,omitempty"`
	MsgType     string             `bson:"msgType,omitempty"`
	MsgContType string             `bson:"msgContType,omitempty"`
	State       string             `bson:"state,omitempty"`
}
