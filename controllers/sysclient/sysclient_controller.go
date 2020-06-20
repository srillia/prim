package sysclient

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"prim/common"
	"prim/controllers"
	"prim/lib/mongolib"
	"prim/lib/redislib"
	"prim/models"
)

//todo 做系统用户的登录和授权

//根据第三方系统用户账户，获取系统客户端
func GetSysclientByAccount(c *gin.Context) {
	account := c.Query("account")
	data := make(map[string]interface{})
	sysclient := getSysclient(account)
	data["sysclient"] = sysclient
	controllers.Response(c, common.OK, "", data)
}

func getSysclient(account string) interface{} {
	return mongolib.FindOne(mongolib.GetConn("prim_sysclient"), bson.M{"account": account}, models.PrimSysClient{})
}

// 查看全部在线用户
func CreateSysclientByPhoneNum(c *gin.Context) {
	phoneNum := c.Query("phoneNum")
	account := c.Query("account")
	password := c.Query("password")
	data := make(map[string]interface{})
	createSysclient(phoneNum, account, password)
	controllers.Response(c, common.OK, "", data)
}

func createSysclient(phoneNum string, account string, password string) {

	//Md5加密
	md5Password := common.EncryptByMd5(password)
	sysclient := models.PrimSysClient{}
	sysclient.Account = account
	sysclient.Password = md5Password
	sysclient.PhoneNum = phoneNum
	sysclient.State = 1
	//生成AuthCode
	sysclient.AuthCode = common.GenerateAuthCode(account)

	mongolib.InsertOne(mongolib.GetConn("prim_sysclient"), sysclient)
}

//获取临时用户key，临时用户key，是用来，标记第三方系统的用户具有连接prim的权限的身份
func GetTempUserKeyBySysclientAuthCode(c *gin.Context) {
	sysAccount := c.Query("sysAccount")
	authCode := c.Query("authCode")
	userId := c.Query("userId")
	data := make(map[string]interface{})
	if checkAuthCode(sysAccount, authCode) {

		tempKey := generateTempKey(sysAccount, userId)
		redislib.SaveTempKey(sysAccount, userId, tempKey)
		data["tempKey"] = tempKey
		controllers.Response(c, common.OK, "", data)
	} else {
		controllers.Response(c, common.Unauthorized, "", data)
	}
}

func generateTempKey(account string, userId string) string {
	return common.EncryptByMd5(account + "：" + userId)
}

func checkAuthCode(account string, code string) bool {
	sysclient := models.PrimSysClient{Account: account, AuthCode: code}
	one := mongolib.FindOne(mongolib.GetConn("prim_sysclient"), sysclient, models.PrimSysClient{})
	if one != nil {
		return true
	} else {
		return false
	}

}
