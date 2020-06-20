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
	client *redis.Client
)

func ExampleNewClient() {

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

func SaveTempKey(sysAccount string, userId string, tempKey string) {
	//hash key 是由 sysAccount加用户id组成
	hkey := sysAccount + ":tempKey"

	number, err := client.Do("hSet", hkey, userId, tempKey).Int()
	if err != nil {
		fmt.Println("saveTempKeyInRedis", hkey, number, err)
		return
	}
}

func CheckTempKey(sysAccount string, userId string) bool {
	//hash key 是由 sysAccount加用户id组成
	hkey := sysAccount + ":tempKey"
	requestTempKey := common.GenerateTempKey(sysAccount, userId)
	tempKey, err := client.Do("hGet", hkey, userId).String()
	if err != nil {
		if requestTempKey == tempKey {
			return true
		}
	}
	return false
}
