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
	key := "rank"
	// 给Golang + 10分
	newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang").Result()
	fmt.Printf("Golang's score is %f now.\n", newScore)
	newScore, err = redisdb.ZIncrBy(key, 2.0, "PHP").Result()
	fmt.Printf("Golang's score is %f now.\n", newScore)
	newScore, err = redisdb.ZIncrBy(key, 8.0, "Python").Result()
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 按照分值排序
	vids, _ := redisdb.ZRevRange(key, 0, 9).Result()
	fmt.Printf("Golang's score is %s now.\n", strings.Join(vids, ","))

	// 删除key
	redisdb.Del(key).Err()
}
