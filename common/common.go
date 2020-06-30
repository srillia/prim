package common

import (
	"crypto/md5"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"reflect"
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

func CopyProperties(src, dst interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针，.Elem()类似于*ptr的操作返回指针指向的地址反射类型
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil

}
