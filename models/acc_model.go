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

//1对1聊天
type Msg struct {
	SysAccount  string `json:"sysAccount,omitempty"`
	SenderId    string `json:"senderId,omitempty"`
	ReceiverId  string `json:"receiverId,omitempty"`
	Time        int64  `json:"time,omitempty"`
	DateTime    string `json:"dateTime,omitempty"`
	Message     string `json:"message,omitempty"`
	RoomId      string `json:"roomId,omitempty"`
	MsgType     string `json:"msgType,omitempty"`     // info 信息，sysMsg系统消息，robot机器人
	MsgContType string `json:"msgContType,omitempty"` // text文本，emoj 表情,file文件，picture 图片，audio 语音，video视频
}

type GroupMsg struct {
	//todo
}

// 心跳请求数据
type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}
