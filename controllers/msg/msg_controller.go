package msg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"prim/common"
	"prim/controllers"
	"prim/lib/redislib"
	"prim/models"
	"prim/servers/websocket"
	"strconv"
	"time"
)

//todo 做系统用户的登录和授权

//根据第三方系统用户账户，获取系统客户端
func SendMsg(c *gin.Context) {
	data := make(map[string]interface{})
	token := c.Query("token")

	pass, sysAccount, appPlatform, userId := redislib.PassCheckToken(token)
	if pass == false {
		//直接return,没有权限连接
		controllers.Response(c, common.Unauthorized, "", data)
		return
	}

	//模拟生成一个client，只有三个属性
	client := &websocket.Client{}
	client.SysAccount = sysAccount
	client.AppPlatform = appPlatform
	client.UserId = userId

	msg := &models.Msg{}
	msg.ReceiverId = c.PostForm("receiverId")
	msg.Message = c.PostForm("message")
	timeInt64, err := strconv.ParseInt(c.PostForm("time"), 10, 64)
	if err == nil {
		msg.Time = timeInt64
	} else {
		msg.Time = time.Now().UnixNano()
	}
	msg.SenderId = c.PostForm("senderId")
	msg.RoomId = c.PostForm("roomId")
	msg.MsgContType = c.PostForm("msgContType")
	msg.MsgType = c.PostForm("msgType")

	acc := &models.Acc{}
	//通过http发送信息，没有acc唯一id
	acc.Seq = ""
	acc.Action = "msg"
	acc.Msg = &msg

	websocket.SaveMessageInMongo(client, acc)
	fmt.Printf("给用户发送消息:%v", msg)
	websocket.SendMessageToReceivers(client, acc, []string{msg.ReceiverId})

	controllers.Response(c, common.OK, "", data)
}
