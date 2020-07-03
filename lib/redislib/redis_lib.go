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

func SaveToken(tokenField string, sysAccount string, appPlatform string, userId string, avatar string, nickname string) {
	key := addPrefix(tokenField)

	client.Do("hSet", key, "sysAccount", sysAccount).Int()
	client.Do("hSet", key, "appPlatform", appPlatform).Int()
	client.Do("hSet", key, "userId", userId).Int()
	client.Do("hSet", key, "avatar", avatar).Int()
	client.Do("hSet", key, "nickname", nickname).Int()

	client.Do("expire", key, 12*3600)
}

func addPrefix(tokenField string) (token string) {
	return primTokeyKey + ":" + tokenField
}

func PassCheckToken(token string) (isPass bool, sysAccount string, appPlatform string, userId string, avatar string, nickname string) {
	key := addPrefix(token)

	_, err := client.Exists(key).Result()
	//认证通过
	if err == nil {
		isPass = true
		sysAccount = client.HGet(key, "sysAccount").Val()
		appPlatform = client.HGet(key, "appPlatform").Val()
		userId = client.HGet(key, "userId").Val()
		avatar = client.HGet(key, "avatar").Val()
		nickname = client.HGet(key, "nickname").Val()
	} else {
		isPass = false
	}
	return
}
