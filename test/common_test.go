package test

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"os/exec"
	"prim/common"
	"prim/initialize"
	"prim/models"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	dir, _ := os.Getwd()
	index := strings.LastIndex(dir, "\\")
	initialize.SetConfigPath(dir[:index])
	initialize.InitConfig()
	m.Run()
}

func TestUUID(t *testing.T) {
	// 创建
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}

func TestChan(t *testing.T) {

	conn := make(chan int, 0)
	go func() {
		conn <- 20
	}()

	message, ok := <-conn
	fmt.Printf("Successfully parsed: %v,%v", message, ok)

}

func TestPath2(t *testing.T) {

}

func TestPath(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	index := strings.LastIndex(dir, "\\")
	var configPath string
	// 分割符是以\\开头
	if index != -1 {
		configPath = dir[:index] + "\\config"

	} else { // 分割符是以/开头
		index := strings.LastIndex(dir, "/")
		configPath = dir[:index] + "/config"
	}

	fmt.Println(configPath)
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func TestData(t *testing.T) {
	message := "{\"seq\":\"461349684816\",\"action\":\"msg\",\"msg\":{\"sysAccount\":\"tokey\",\"senderId\":\"9527\",\"receiverId\":\"4399\",\"time\":\"1592893964624\",\"message\":\"4399你好，我是9527\",\"msyType\":\"info\",\"msyContType\":\"text\"}}"
	//msg := "{\"seq\":\"461349684816\"}"
	acc := &models.Acc{}
	err := json.Unmarshal([]byte(message), acc)
	fmt.Printf("%v", acc)
	fmt.Println(err)
}

func TestSplit(t *testing.T) {
	arr := strings.Split("xxx_sss_", "_")
	fmt.Printf("%v", arr)
	fmt.Printf("%d****%s****\n", len(arr), arr[2])
}

func TestToken(t *testing.T) {
	tokenField := common.EncryptByMd5("tokey_web_9527" + strconv.FormatInt(time.Now().UnixNano(), 10))
	fmt.Printf("%v\n", time.Now().UnixNano())
	string := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Printf("%v\n", string)
	fmt.Printf("%v\n", tokenField)

	tokenField2 := common.EncryptByMd5("tokey_web_9527")
	fmt.Printf("%v\n", tokenField2)
}
func TestTime(t *testing.T) {
	datetime := "2015-01-01 00:00:00" //待转化为时间戳的字符串

	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	fmt.Println(timestamp)

	//时间戳转化为日期
	datetime = time.Unix(1592961359393/1000, 1592961359393%1000).Format(timeLayout)
	fmt.Println(datetime)

}

func TestNum(t *testing.T) {
	time := 1592961359393

	time2 := time % 1000
	time3 := time / 1000

	fmt.Printf("%v\n", time2)
	fmt.Printf("%v\n", time3)

}
