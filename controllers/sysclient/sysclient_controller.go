package sysClient

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"prim/common"
	"prim/controllers"
	"prim/lib/mongolib"
	"prim/lib/redislib"
	"prim/models"
	"prim/servers/websocket"
)

//todo 做系统用户的登录和授权

//根据第三方系统用户账户，获取系统客户端
func GetSysClient(c *gin.Context) {
	account := c.Query("account")
	data := make(map[string]interface{})
	sysClient, err := getSysClient(account)
	if err != nil {
		controllers.Response(c, common.NotData, "", data)
		return
	}
	data["sysClient"] = sysClient
	controllers.Response(c, common.OK, "", data)
}

func getSysClient(account string) (interface{}, error) {
	sysClient := &models.PrimSysClient{}
	err := mongolib.FindOne(mongolib.GetConn("prim_sys_client"), bson.M{"account": account}, sysClient)
	return sysClient, err
}

// 查看全部在线用户
func CreateSysClient(c *gin.Context) {
	phoneNum := c.PostForm("phoneNum")
	account := c.PostForm("account")
	password := c.PostForm("password")
	data := make(map[string]interface{})

	err := saveSysClient(phoneNum, account, password)
	if err == nil {
		controllers.Response(c, common.OK, "", data)
	} else {
		controllers.Response(c, common.SysAccountExist, "", data)
	}

}

func saveSysClient(phoneNum string, account string, password string) error {

	client := &models.PrimSysClient{}
	err := mongolib.FindOne(mongolib.GetConn("prim_sys_client"),
		bson.D{{"account", account}}, client)
	fmt.Printf("%v", client)

	//说明已经存在客户端
	if err == nil {
		return errors.New("系统用户:账号" + account + ",已经存在")

	}

	//Md5加密
	md5Password := common.EncryptByMd5(password)
	sysClient := models.PrimSysClient{}
	sysClient.Account = account
	sysClient.Password = md5Password
	sysClient.PhoneNum = phoneNum
	sysClient.State = 1
	//生成AuthCode
	sysClient.AuthCode = common.GenerateAuthCode(account)

	mongolib.InsertOne(mongolib.GetConn("prim_sys_client"), sysClient)
	return nil
}

//获取临时用户key，临时用户key，是用来，标记第三方系统的用户具有连接prim的权限的身份
func GetToken(c *gin.Context) {
	authCode := c.Query("authCode")
	sysAccount := c.Query("sysAccount")
	appPlatform := c.Query("appPlatform")
	userId := c.Query("userId")
	avatar := c.Query("avatar")
	nickname := c.Query("nickname")
	data := make(map[string]interface{})
	if checkAuthCode(sysAccount, authCode) {

		tokeyKey := common.GenerateTokenKey(sysAccount, appPlatform, userId)
		redislib.SaveToken(tokeyKey, sysAccount, appPlatform, userId, avatar, nickname)
		data["token"] = tokeyKey
		controllers.Response(c, common.OK, "", data)
	} else {
		controllers.Response(c, common.Unauthorized, "", data)
	}
}

func checkAuthCode(account string, code string) bool {
	primSysClient := &models.PrimSysClient{}
	err := mongolib.FindOne(mongolib.GetConn("prim_sys_client"), bson.D{{"account", account}, {"authCode", code}}, primSysClient)
	//err == nil 说明存在
	if err == nil {
		return true
	} else {
		return false
	}
}

//退出连接
func ExitClient(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		controllers.Response(c, common.ParameterIllegal, "", nil)
	}
	pass, sysAccount, appPlatform, userId, _, _ := redislib.PassCheckToken(token)
	if pass == false {
		//直接return,没有权限连接
		controllers.Response(c, common.Unauthorized, "", nil)
		return
	}
	client := websocket.GetUserClient(sysAccount, appPlatform, userId)
	acc := &models.Acc{}
	acc.Seq = ""
	acc.Action = "msg"
	acc.Msg = nil
	websocket.ClearClient(client, acc)
	controllers.Response(c, common.OK, "", nil)
}
