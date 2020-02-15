package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()
}

func main() {
	Init()
	fmt.Println(os.Getenv("OSS_END_POINT"))
}
