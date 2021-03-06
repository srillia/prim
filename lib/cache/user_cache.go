/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 17:28
 */

package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"prim/lib/redislib"
	"prim/models"
)

const (
	userOnlinePrefix    = "acc:user:online:" // 用户在线状态
	userOnlineCacheTime = 24 * 60 * 60
)

/*********************  查询用户是否在线  ************************/
func getUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)
	return
}

func GetUserOnlineInfo(userKey string) (userOnline *models.User, err error) {
	redisClient := redislib.GetClient()

	key := getUserOnlineKey(userKey)

	data, err := redisClient.Get(key).Bytes()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("GetUserOnlineInfo", userKey, err)

			return
		}

		fmt.Println("GetUserOnlineInfo", userKey, err)

		return
	}

	userOnline = &models.User{}
	err = json.Unmarshal(data, userOnline)
	if err != nil {
		fmt.Println("获取用户在线数据 json Unmarshal", userKey, err)

		return
	}

	fmt.Println("获取用户在线数据", userKey, "time", userOnline.LoginTime, userOnline.HeartbeatTime, "AccIp", userOnline.AccIp, userOnline.IsLogoff)

	return
}

//TODO 往redis中存用户的数据
//设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *models.User) (err error) {

	redisClient := redislib.GetClient()
	key := getUserOnlineKey(userKey)

	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		fmt.Println("设置用户在线数据 json Marshal", key, err)

		return
	}

	_, err = redisClient.Set(key, string(valueByte), -1).Result()
	if err != nil {
		fmt.Println("设置用户在线数据 ", key, err)

		return
	}
	return
}
