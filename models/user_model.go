/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 17:36
 */

package models

import (
	"fmt"
	"time"
)

const (
	heartbeatTimeout = 3 * 60 // 用户心跳超时时间
)

// 用户在线状态
type User struct {
	AccIp         string `json:"accIp"`         // acc Ip
	AccPort       string `json:"accPort"`       // acc 端口
	AppPlatform   string `json:"appPlatform"`   // appPlatform
	SysAccount    string `json:"sysAccount"`    // appPlatform
	UserId        string `json:"userId"`        // 用户Id
	Avatar        string `json:"avatar"`        // avatar 头像
	Nickname      string `json:"nickname"`      // nickname 昵称
	ClientIp      string `json:"clientIp"`      // 客户端Ip
	ClientPort    string `json:"clientPort"`    // 客户端端口
	LoginTime     uint64 `json:"loginTime"`     // 用户上次登录时间
	HeartbeatTime uint64 `json:"heartbeatTime"` // 用户上次心跳时间
	LogOutTime    uint64 `json:"logOutTime"`    // 用户退出登录的时间
	Qua           string `json:"qua"`           // qua
	DeviceInfo    string `json:"deviceInfo"`    // 设备信息
	IsLogoff      bool   `json:"isLogoff"`      // 是否下线
}

/**********************  数据处理  *********************************/

// 用户登录
func NewUser(accIp, accPort string, appPlatform string, sysAccount string, userId string, avatar string, nickname string, addr string, loginTime uint64) (userOnline *User) {

	userOnline = &User{
		AccIp:         accIp,
		AccPort:       accPort,
		AppPlatform:   appPlatform,
		SysAccount:    sysAccount,
		UserId:        userId,
		Avatar:        avatar,
		Nickname:      nickname,
		ClientIp:      addr,
		LoginTime:     loginTime,
		HeartbeatTime: loginTime,
		IsLogoff:      false,
	}

	return
}

// 用户心跳
func (u *User) Heartbeat(currentTime uint64) {

	u.HeartbeatTime = currentTime
	u.IsLogoff = false

	return
}

// 用户退出登录
func (u *User) LogOut() {

	currentTime := uint64(time.Now().Unix())
	u.LogOutTime = currentTime
	u.IsLogoff = true

	return
}

/**********************  数据操作  *********************************/

// 用户是否在线
func (u *User) IsOnline() (online bool) {
	if u.IsLogoff {

		return
	}

	currentTime := uint64(time.Now().Unix())

	if u.HeartbeatTime < (currentTime - heartbeatTimeout) {
		fmt.Println("用户是否在线 心跳超时", u.AppPlatform, u.UserId, u.HeartbeatTime)

		return
	}

	if u.IsLogoff {
		fmt.Println("用户是否在线 用户已经下线", u.AppPlatform, u.UserId)

		return
	}

	return true
}

// 用户是否在本台机器上
func (u *User) UserIsLocal(localIp, localPort string) (result bool) {

	if u.AccIp == localIp && u.AccPort == localPort {
		result = true

		return
	}

	return
}
