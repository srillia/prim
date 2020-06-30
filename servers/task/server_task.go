/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-03
* Time: 15:44
 */

package task

import (
	"fmt"
	"prim/initialize"
	"prim/lib/cache"
	"runtime/debug"
	"time"
)

func GrpcServerRegisterInit() {
	Timer(5*time.Millisecond, 60*time.Second, server, "", serverDefer, "")
}

// 服务注册
func server(param interface{}) (result bool) {
	result = true

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务注册 stop", r, string(debug.Stack()))
		}
	}()

	server := initialize.GetServer()
	currentTime := uint64(time.Now().Unix())
	fmt.Println("定时任务，服务注册", param, server, currentTime)

	if server.Ip != "" && server.Port != "" {
		cache.SetServerInfo(server, currentTime)
	}
	return
}

// 服务下线
func serverDefer(param interface{}) (result bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务下线 stop", r, string(debug.Stack()))
		}
	}()

	fmt.Println("服务下线", param)

	server := initialize.GetServer()
	cache.DelServerInfo(server)

	return
}
