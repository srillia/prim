package models

/************************  请求数据  **************************/
// 通用websocket acc请求数据格式
type Acc struct {
	//客户端生成的消息唯一id，用于信息重发时的，去重工作
	Seq string `json:"seq"` // 消息的唯一Id
	//Action是处理信息的行为路由
	Action string `json:"action"` // 请求命令字
	//信息的数据包
	Msg interface{} `json:"msg,omitempty"` // 消息体
}

func (acc *Acc) OkAcc(msg interface{}) *Acc {
	//需要根据acc获取seq
	return &Acc{acc.Seq, "ok", msg}
}

func (acc *Acc) AckAcc(msg interface{}) *Acc {
	//需要根据acc获取seq
	return &Acc{acc.Seq, "ack", msg}
}

func (acc *Acc) HeartBeatAcc(msg interface{}) *Acc {
	//需要根据acc获取seq
	return &Acc{acc.Seq, "heartbeat", msg}
}
