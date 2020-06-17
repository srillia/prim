package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PrimSysClient struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Account  string             `bson:"account,omitempty"`
	Password string             `bson:"password,omitempty"`
	State    byte               `bson:"state,omitempty"`
	authCode string
}

type PrimUser struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	UserId          string
	PrimSysClientId primitive.ObjectID
	Account         string
	Nickname        string
	Signature       string
	Sex             int8
	Birthday        time.Time
	Tel             string
	Email           string
	Intro           string
	Avatar          string
	Age             int8
	State           byte
}

type PrimRoom struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserList []string
}

type PrimMessage struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Message     string
	RoomId      primitive.ObjectID
	SenderId    string
	ReceiverId  string
	MsgType     byte
	MsgContType byte
	State       byte
}
