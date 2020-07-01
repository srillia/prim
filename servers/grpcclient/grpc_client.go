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
func SendMsg(server *models.Server, sysAccount string, appPlatform string, receiverIds []string, acc []byte) (err error) {
	fmt.Println("Grpc被调用 " + common.GetFuncName() + "-------------")
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println(common.GetFuncName(), server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.SendMsgReq{
		SysAccount:  sysAccount,
		AppPlatform: appPlatform,
		UserIds:     receiverIds,
		Acc:         acc,
	}
	rsp, err := c.SendMsg(ctx, &req)
	if err != nil {
		fmt.Println(common.GetFuncName(), err)

		return
	}

	if rsp.GetRspCode() != common.OK {
		fmt.Println("SendMsg", rsp.String())
		return
	}

	return
}

func ClearExistsClient(server *models.Server, sysAccount string, appPlatform string, userId string) (err error) {

	fmt.Println("Grpc被调用 " + common.GetFuncName() + "-------------")
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.ClearClientReq{
		SysAccount:  sysAccount,
		AppPlatform: appPlatform,
		UserId:      userId,
	}
	rsp, err := c.ClearExistsClient(ctx, &req)
	if err != nil {
		fmt.Println(common.GetFuncName(), err)
		return
	}

	if rsp.GetRspCode() != common.OK {
		fmt.Println(common.GetFuncName(), rsp.String())

		return
	}
	return
}

// 获取用户列表
func GetUserList(server *models.Server) (userIds []string, err error) {
	userIds = make([]string, 0)

	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println(common.GetFuncName(), server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) //超时时间，太短，将存在访问响应会丢失
	defer cancel()

	req := protobuf.GetUserListReq{}
	rsp, err := c.GetUserList(ctx, &req)
	if err != nil {
		fmt.Println(common.GetFuncName(), err)

		return
	}

	if rsp.GetRspCode() != common.OK {
		fmt.Println("SendMsg", rsp.String())
		return
	}
	userIds = rsp.GetUserIds()
	fmt.Println(common.GetFuncName(), userIds)

	return
}
