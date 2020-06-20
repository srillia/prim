package common

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
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
