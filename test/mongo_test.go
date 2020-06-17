package test

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"prim/lib/mongolib"
	"prim/models"
	"testing"
)

func TestInsert(t *testing.T) {

	t.Log("hello world")
	sysClient := models.PrimSysClient{Account: "1353358523733", Password: "123456", State: 1}

	mongolib.InsertOne(mongolib.GetConn("prim_message"), sysClient)
}

func TestGetOne(t *testing.T) {

	t.Log("hello world")

	sysUser := mongolib.FindOne(mongolib.GetConn("prim_message"), bson.M{"account": "13533585237"}, &models.PrimSysClient{})
	fmt.Printf("%v", sysUser)
}
