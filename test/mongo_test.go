package test

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"prim/lib/mongolib"
	"prim/models"
	"testing"
)

func TestInsert(t *testing.T) {

	t.Log("hello world")
	sysClient := &models.PrimMessage{SysAccount: "1353358523733"}

	mongolib.InsertOne(mongolib.GetConn("prim_message"), sysClient)
}

func TestGetOne(t *testing.T) {

	//t.Log("hello world")
	//
	//sysUser, _ := mongolib.FindOne(mongolib.GetConn("prim_message"), bson.M{"account": "13533585237"}, models.PrimSysClient{})
	//fmt.Printf("%v", sysUser)
	userList := [2]string{"9532", "9531"}
	primRoom := &models.PrimRoom{}
	err := mongolib.FindOne(mongolib.GetConn("prim_room"), bson.D{{"userList", bson.D{{"$all", userList}}}}, primRoom)
	fmt.Println(primRoom, err)

}

func TestExist(t *testing.T) {

	t.Log("hello world")
	id, _ := primitive.ObjectIDFromHex("5ee9b032003953889c5cf6dd")
	sysClient := &models.PrimSysClient{}
	fmt.Printf("%v\n", sysClient)
	SingleResult := mongolib.GetConn("prim_message").FindOne(mongolib.GetContext(), bson.D{{"_id", id}})
	err := SingleResult.Decode(sysClient)

	//err为nil说明存在记录
	if err == nil {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}

func TestPrimRoom(t *testing.T) {

	arr := [2]string{"789794", "123456789"}
	primRoom := models.PrimRoom{UserList: arr}

	mongolib.InsertOne(mongolib.GetConn("prim_room"), primRoom)

}

func TestFindPrimRoom(t *testing.T) {

	arr := [2]string{"9532", "9531"}
	//primRoom := models.PrimRoom{UserList:arr}
	room := &models.PrimRoom{}
	err := mongolib.FindOne(mongolib.GetConn("prim_room"), bson.D{{"userList", bson.D{{"$all", arr}}}}, room)
	fmt.Println(room.Id, err)

}
