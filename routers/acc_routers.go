/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:02
 */

package routers

import (
	"prim/servers/websocket"
)

// Websocket 路由
func WebsocketInit() {
	websocket.Register("heartbeat", websocket.HeartbeatController)
	websocket.Register("exit", websocket.ExitController)
	websocket.Register("ping", websocket.PingController)
	websocket.Register("msg", websocket.MsgController)
}
