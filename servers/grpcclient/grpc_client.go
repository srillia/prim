/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-03
* Time: 16:43
 */

package grpcclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"prim/common"
	"prim/models"
	"prim/protobuf"
	"time"
)

// rpc 发送一对一信息
func SendMsg(server *models.Server, seq string, sysAccount string, appPlatform string, receiverIds []string, message []byte) (sendMsgId string, err error) {
	fmt.Println("查看是否调用了grpc_client.go。SendMsgAll-------------")
	// Set up a connection to the server.
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.SendMsgReq{
		Seq:         seq,
		SysAccount:  sysAccount,
		AppPlatform: appPlatform,
		UserIds:     receiverIds,
		Msg:         message,
	}
	rsp, err := c.SendMsg(ctx, &req)
	if err != nil {
		fmt.Println("给全体用户发送消息", err)

		return
	}

	if rsp.GetRetCode() != common.OK {
		fmt.Println("给全体用户发送消息", rsp.String())

		return
	}

	sendMsgId = rsp.GetSendMsgId()
	fmt.Println("给全体用户发送消息 成功:", sendMsgId)

	return
}

// rpc client
// 给全体用户发送消息
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func SendMsgAll(server *models.Server, seq string, appPlatform string, userId string, action string, message string) (sendMsgId string, err error) {
	fmt.Println("查看是否调用了grpc_client.go。SendMsgAll-------------")
	// Set up a connection to the server.
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.SendMsgAllReq{
		Seq:         seq,
		AppPlatform: appPlatform,
		UserId:      userId,
		Action:      action,
		Msg:         message,
	}
	rsp, err := c.SendMsgAll(ctx, &req)
	if err != nil {
		fmt.Println("给全体用户发送消息", err)

		return
	}

	if rsp.GetRetCode() != common.OK {
		fmt.Println("给全体用户发送消息", rsp.String())

		return
	}

	sendMsgId = rsp.GetSendMsgId()
	fmt.Println("给全体用户发送消息 成功:", sendMsgId)

	return
}

// 获取用户列表
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func GetUserList(server *models.Server) (userIds []string, err error) {
	userIds = make([]string, 0)

	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)

	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	req := protobuf.GetUserListReq{}
	rsp, err := c.GetUserList(ctx, &req)
	if err != nil {
		fmt.Println("获取用户列表 发送请求错误:", err)

		return
	}

	if rsp.GetRetCode() != common.OK {
		fmt.Println("获取用户列表 返回码错误:", rsp.String())

		return
	}

	userIds = rsp.GetUserIds()
	fmt.Println("获取用户列表 成功:", userIds)

	return
}
