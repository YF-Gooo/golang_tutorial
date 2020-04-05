package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()
}
func main() {
	Init()
	fmt.Println()
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/assets", "./assets")
	router.StaticFile("/", "./assets/index.html")

	router.GET("/connect", func(c *gin.Context) {
		url := fmt.Sprintf("https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=%s", os.Getenv("APPID"), os.Getenv("REDIRECT_URI"), os.Getenv("STATE"))
		c.Redirect(http.StatusMovedPermanently, url)
	})
	router.GET("/auth", func(c *gin.Context) {
		fmt.Print(c.Query("code"))
		fmt.Print(c.Query("state"))
		url := fmt.Sprintf("https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=[YOUR_APP_ID]&client_secret=[YOUR_APP_Key]&code=[The_AUTHORIZATION_CODE]&redirect_uri=[YOUR_REDIRECT_URI]")
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(body)

	})
	router.Run(":3000")
}
