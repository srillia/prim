/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-03
* Time: 16:43
 */

package grpcserver

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"prim/helper"
	"prim/protobuf"
	"prim/servers/websocket"
)

type server struct {
}

//func setErr(rsp proto.Message, code uint32, message string) {
//
//	message = common.GetErrorMessage(code, message)
//	switch v := rsp.(type) {
//	case *protobuf.QueryUsersOnlineRsp:
//		v.RetCode = code
//		v.ErrMsg = message
//	case *protobuf.SendMsgRsp:
//		v.RetCode = code
//		v.ErrMsg = message
//	case *protobuf.SendMsgAllRsp:
//		v.RetCode = code
//		v.ErrMsg = message
//	case *protobuf.GetUserListRsp:
//		v.RetCode = code
//		v.ErrMsg = message
//	default:
//
//	}
//
//}

// 查询用户是否在线
func (s *server) QueryUsersOnline(c context.Context, req *protobuf.QueryUsersOnlineReq) (rsp *protobuf.QueryUsersOnlineRsp, err error) {

	fmt.Println("grpc_request 查询用户是否在线", req.String())

	rsp = &protobuf.QueryUsersOnlineRsp{}

	online := websocket.CheckUserOnline(req.GetSysAccount(), req.GetAppPlatform(), req.GetUserId())

	//setErr(req, common.OK, "")
	rsp.Online = online

	return rsp, nil
}

// 给本机用户发消息
func (s *server) SendMsg(c context.Context, req *protobuf.SendMsgReq) (rsp *protobuf.SendMsgRsp, err error) {

	fmt.Println("grpc_request 给本机用户发消息", req.String())

	websocket.SendUserMessageLocal(req.GetSysAccount(), req.GetAppPlatform(), req.GetUserIds(), req.Acc)
	return
}

////给本机全体用户发消息
//func (s *server) SendMsgAll(c context.Context, req *protobuf.SendMsgAllReq) (rsp *protobuf.SendMsgAllRsp, err error) {
//
//	fmt.Println("grpc_request 给本机全体用户发消息", req.String())
//
//	rsp = &protobuf.SendMsgAllRsp{}
//
//	data := models.GetMsgData(req.GetUserId(), req.GetSeq(), req.GetAction(), req.GetMsg())
//	websocket.AllSendMessages(req.GetSysAccount(), req.GetAppPlatform(), req.GetUserId(), data)
//
//	setErr(rsp, common.OK, "")
//
//	fmt.Println("grpc_response 给本机全体用户发消息:", rsp.String())
//
//	return
//}

// 获取本机用户列表
func (s *server) GetUserList(c context.Context, req *protobuf.GetUserListReq) (rsp *protobuf.GetUserListRsp, err error) {

	fmt.Println("grpc_request 获取本机用户列表", req.String())

	rsp = &protobuf.GetUserListRsp{}

	// 本机
	userList := websocket.GetUserList()

	//setErr(rsp, common.OK, "")
	rsp.UserIds = userList

	fmt.Println("grpc_response 获取本机用户列表:", rsp.String())

	return
}

// rpc server
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
func Init() {

	rpcPort := viper.GetString("app.rpcPort")
	lis, err := net.Listen("tcp", ":"+rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterAccServerServer(s, &server{})
	fmt.Println("Grpc Server 启动成功", helper.GetServerIp(), rpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
