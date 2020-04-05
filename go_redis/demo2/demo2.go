package main

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

// redis

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return
	}
	fmt.Println("连接redis成功！")
	// zset
	key := "ranklist"
	redisdb.LPush(key, 10.0)
	redisdb.LPush(key, 9.0)
	redisdb.LPush(key, 8.0)
	redisdb.LPush(key, 7.0)
	redisdb.LPush(key, 3.0)
	x, _ := redisdb.RPop(key).Result()
	fmt.Printf("POP's NUMBER is %f now.\n", x)
	x, _ = redisdb.RPop(key).Result()
	fmt.Printf("POP's NUMBERis %f now.\n", x)
	x, _ = redisdb.RPop(key).Result()
	fmt.Printf("POP's NUMBER is %f now.\n", x)
	y, _ := redisdb.LRange(key, 0, 10).Result()
	fmt.Printf(strings.Join(y, ","))
}
