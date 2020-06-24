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
	"prim/common"
	"prim/models"
	"sync"
)

type DisposeFunc func(client *Client, seq string, message []byte) (code uint32, msg string, data interface{})

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

	requestData, err := json.Marshal(acc.Msg)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)
		client.SendMsg([]byte("处理数据失败"))

		return
	}

	seq := acc.Seq
	action := acc.Action

	var (
		code uint32
		msg  string
		data interface{}
	)

	// request
	fmt.Println("acc_request", action, client.Addr)

	// 采用 map 注册的方式
	if handle, ok := getHandlers(action); ok {
		code, msg, data = handle(client, seq, requestData)
	} else {
		code = common.RoutingNotExist
		fmt.Println("处理数据 路由不存在", client.Addr, "action", action)
	}

	msg = common.GetErrorMessage(code, msg)

	responseHead := models.NewResponseHead(seq, action, code, msg, data)

	headByte, err := json.Marshal(responseHead)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)

		return
	}

	client.SendMsg(headByte)

	fmt.Println("acc_response send", client.Addr, client.AppPlatform, client.UserId, "action", action, "code", code)

	return
}
