/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 14:18
 */

package redislib

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"prim/common"
)

var (
	client       *redis.Client
	primTokeyKey = "primTokenKey"
)

func NewClient() {

	client = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	pong, err := client.Ping().Result()
	fmt.Println("初始化redis:", pong, err)
	// Output: PONG <nil>
}

func GetClient() (c *redis.Client) {

	return client
}

func SaveToken(tokenField string, tokenValue string) {
	client.Do("set", addPrefix(tokenField), tokenValue)
	client.Do("expire", addPrefix(tokenField), 12*3600)
}

func addPrefix(tokenField string) (token string) {
	return primTokeyKey + ":" + tokenField
}

func PassCheckToken(token string) (isPass bool, sysAccount string, appPlatform string, userId string) {
	//处理token
	tokenValue, err := client.Do("get", addPrefix(token)).String()
	arr := common.ParseToken(tokenValue)
	//err==有nil说明有值
	if err == nil {

		//todo grpc访问所有的服务器，删除已经连接的请求
		return true, arr[0], arr[1], arr[2]
	}
	return false, "", "", ""
}
