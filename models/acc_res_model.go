package models

//1对1聊天
type Ack struct {
	Result string `json:"result,omitempty"` // ack,no-ack
}

type Ok struct {
	Result string `json:"result,omitempty"`
}

// 心跳请求数据
type HeartBeat struct {
	Result string `json:"result,omitempty"`
}
