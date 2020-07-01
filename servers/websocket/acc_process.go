/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 14:38
 */

package websocket

import (
	"encoding/json"
	"fmt"
	"prim/models"
	"sync"
)

type DisposeFunc func(client *Client, acc *models.Acc) *models.Acc

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
)

// 注册
func Register(key string, value DisposeFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value

	return
}

func getHandlers(key string) (handle DisposeFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()

	handle, ok = handlers[key]

	return
}

// 处理数据
func ProcessData(client *Client, message []byte) {

	fmt.Println("处理数据", client.Addr, string(message))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据 stop", r)
		}
	}()

	acc := &models.Acc{}

	err := json.Unmarshal(message, acc)
	if err != nil {
		fmt.Println("处理数据 json Unmarshal", err)
		client.SendMsg([]byte("数据不合法"))

		return
	}

	//requestData, err := json.Marshal(acc.Msg)
	//if err != nil {
	//	fmt.Println("处理数据 json Marshal", err)
	//	client.SendMsg([]byte("处理数据失败"))
	//
	//	return
	//}

	var (
		ret *models.Acc
	)

	// request
	fmt.Println("acc_request", acc.Action, client.Addr)
	// 采用 map 注册的方式
	if handle, ok := getHandlers(acc.Action); ok {
		ret = handle(client, acc)
	} else {
		//todo 添加 ex action
		fmt.Println("处理数据 路由不存在", client.Addr, "action", acc.Action)
	}

	retByte, err := json.Marshal(ret)
	if err != nil {
		fmt.Println("处理数据 json ret", err)

		return
	}

	//如何action是exit说明client已经退出，连接已经断开
	if ret.Action != "exit" {
		client.SendMsg(retByte)
	}

	fmt.Println("acc_response send", client.Addr, client.AppPlatform, client.UserId, "action", ret.Action)

	return
}
