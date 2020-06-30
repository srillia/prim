package models

//1对1聊天
type Msg struct {
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
