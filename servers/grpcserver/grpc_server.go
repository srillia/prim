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
	"google.golang.org/grpc"
	"log"
	"net"
	"prim/common"
	"prim/initialize"
	"prim/models"
	"prim/protobuf"
	"prim/servers/websocket"
)

type server struct {
}

// 查询用户是否在线
func (s *server) QueryUsersOnline(c context.Context, req *protobuf.QueryUsersOnlineReq) (rsp *protobuf.QueryUsersOnlineRsp, err error) {

	fmt.Println("grpc_request 查询用户是否在线", req.String())

	online := websocket.CheckUserOnline(req.GetSysAccount(), req.GetAppPlatform(), req.GetUserId())

	rsp = &protobuf.QueryUsersOnlineRsp{}
	rsp.RspCode = common.OK
	rsp.Online = online

	return rsp, nil
}

// 给本机用户发消息
func (s *server) SendMsg(c context.Context, req *protobuf.SendMsgReq) (rsp *protobuf.SendMsgRsp, err error) {

	fmt.Println("grpc_request 给本机用户发消息", req.String())
	websocket.SendUserMessageLocal(req.GetSysAccount(), req.GetAppPlatform(), req.GetUserIds(), req.Acc)
	rsp = &protobuf.SendMsgRsp{}
	rsp.RspCode = common.OK
	return
}

// 给本机用户发消息
func (s *server) ClearExistsClient(c context.Context, req *protobuf.ClearClientReq) (rsp *protobuf.ClearClientRsp, err error) {

	fmt.Println("grpc_request 给本机用户发消息", req.String())

	//清除用户在该平台的连接
	client := websocket.GetUserClient(req.GetSysAccount(), req.GetAppPlatform(), req.GetUserId())
	websocket.ClearClient(client, &models.Acc{})

	rsp = &protobuf.ClearClientRsp{}
	rsp.RspCode = common.OK
	return
}

// 获取本机用户列表
func (s *server) GetUserList(c context.Context, req *protobuf.GetUserListReq) (rsp *protobuf.GetUserListRsp, err error) {

	fmt.Println("grpc_request 获取本机用户列表", req.String())
	// 本机
	userList := websocket.GetUserList()

	rsp = &protobuf.GetUserListRsp{}
	rsp.RspCode = common.OK
	rsp.UserIds = userList

	fmt.Println("grpc_response 获取本机用户列表:", rsp.String())
	return
}

// rpc server
func Init() {
	lis, err := net.Listen("tcp", ":"+initialize.RpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterAccServerServer(s, &server{})
	fmt.Println("Grpc Server 启动成功", initialize.ServerIp, initialize.RpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
