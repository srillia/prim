package common

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

func GenerateTempKey(account string, userId string) string {
	return EncryptByMd5(account + "：" + userId)
}

func GenerateAuthCode(account string) string {
	uuid := uuid.NewV4().String()
	authCode := EncryptByMd5(uuid + ":" + account)
	return authCode
}

func EncryptByMd5(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

func GenerateTokenFieldAndValue(account string, appPlatform string, userId string) (tokenField string, tokenValue string) {
	tokenValue = fmt.Sprintf("%s_%s_%s", account, appPlatform, userId)
	tokenField = EncryptByMd5(tokenValue + strconv.FormatInt(time.Now().UnixNano(), 10))
	return
}

func ParseToken(token string) (arr []string) {
	arr = strings.Split(token, "_")
	return
}

func TimestampToDateString(timestamp int64) (dateTime string) {
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	//时间戳转化为日期
	dateTime = time.Unix(timestamp/1000, timestamp%1000).Format(timeLayout)
	return
}
