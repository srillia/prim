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
