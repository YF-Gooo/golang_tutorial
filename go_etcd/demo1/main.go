package main

import (
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"0.0.0.0:12379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	client = client
}
