/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 14:41
 */

package models

//
///************************  请求数据  **************************/
//// 通用请求数据格式
//type Request struct {
//	//客户端生成的消息唯一id，用于信息重发时的，去重工作
//	Seq  string      `json:"seq"`            // 消息的唯一Id
//	//Cmd是处理信息的行为路由
//	Cmd  string      `json:"cmd"`            // 请求命令字
//	//信息的数据包
//	Data interface{} `json:"data,omitempty"` // 数据 json
//}
//
//// 登录请求数据
//type Login struct {
//	ServiceToken string `json:"serviceToken"` // 验证用户是否登录
//	AppPlatform        uint32 `json:"appPlatform,omitempty"`
//	UserId       string `json:"userId,omitempty"`
//}
//
//
//type GroupMsg struct {
//	//todo
//}
//
//// 心跳请求数据
//type HeartBeat struct {
//	UserId string `json:"userId,omitempty"`
//}
