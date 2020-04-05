package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

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

const routineCount = 100

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return
	}
	fmt.Println("连接redis成功！")
	// Transactionally increments key using GET and SET commands.
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			// get current value or zero
			n, err := tx.Get(key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// actual opperation (local in optimistic lock)
			n++

			// runs only if the watched keys remain unchanged
			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				// pipe handles the error case
				pipe.Set(key, n, 0)
				return nil
			})
			return err
		}

		for retries := routineCount; retries > 0; retries-- {
			err := redisdb.Watch(txf, key)
			if err != redis.TxFailedErr {
				return err
			}
			// optimistic lock lost
		}
		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := redisdb.Get("counter3").Int()
	fmt.Println("ended with", n, err)
}
